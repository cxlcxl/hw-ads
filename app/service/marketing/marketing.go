package servicemarketing

import (
	"bs.mobgi.cc/app/model"
	serviceaccount "bs.mobgi.cc/app/service/account"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"errors"
	"fmt"
	"time"
)

type BaseAdsResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type TrackingRequest struct {
	Filtering    TrackingFilter `json:"filtering"`
	AdvertiserId string         `json:"advertiser_id"`
	Page         int            `json:"page"`
	PageSize     int            `json:"page_size"`
}
type TrackingFilter struct {
	ProductUniqueFlag string `json:"product_unique_flag"`
}
type TrackingResponse struct {
	BaseAdsResp
	Data TrackingData `json:"data"`
}
type TrackingData struct {
	Total int64       `json:"total"`
	Data  []*Tracking `json:"data"`
}
type Tracking struct {
	ClickTrackingUrl string `json:"click_tracking_url"`
	TrackingStatus   string `json:"tracking_status"`
	EffectName       string `json:"effect_name"`
	ImpTrackingUrl   string `json:"imp_tracking_url"`
	EffectType       string `json:"effect_type"`
	TrackingId       int64  `json:"tracking_id"`
}
type TrackingItem struct {
	EffectName string `json:"effect_name"`
	EffectType string `json:"effect_type"`
	TrackingId int64  `json:"tracking_id"`
}

func TrackingPull(appId string) (tracks []*TrackingItem, err error) {
	app, err := model.NewApp(vars.DBMysql).FindAppByAppId(appId)
	if err != nil {
		return nil, errors.New("应用信息有误：" + err.Error())
	}
	token, err := serviceaccount.GetToken(app.AccountId)
	if err != nil {
		return nil, errors.New("Token 信息获取失败：" + err.Error())
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return nil, errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	}
	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	page := 1
	d := TrackingRequest{
		Filtering: TrackingFilter{ProductUniqueFlag: app.PkgName},
		Page:      page,
		PageSize:  50,
	}
	items, total, err := pull(t, d, app.AccountId, app.AppId)
	if err != nil {
		return nil, err
	}
	tracks = append(tracks, items...)
	if total > 50 {
		pages := utils.CeilPages(total, 50)
		for i := 2; i <= int(pages); i++ {
			d.Page = i
			if items, _, err = pull(t, d, app.AccountId, app.AppId); err != nil {
				return nil, err
			} else {
				tracks = append(tracks, items...)
			}
		}
	}

	return
}

func pull(token string, d TrackingRequest, accountId int64, appId string) (items []*TrackingItem, total int64, err error) {
	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Tracking.Query")).Get().JsonData(d)
	if err != nil {
		return nil, 0, err
	}
	var response TrackingResponse
	if err = c.Request(&response, curl.Authorization(token)); err != nil {
		return nil, 0, err
	}
	if response.Code != "200" {
		return nil, 0, errors.New("华为接口拉取错误：" + response.Message)
	}
	var trackings []*model.Tracking
	disabled := make([]int64, 0)
	for _, datum := range response.Data.Data {
		if datum.TrackingStatus == vars.TrackingStatusActive {
			trackings = append(trackings, &model.Tracking{
				AppId:            appId,
				AccountId:        accountId,
				EffectType:       datum.EffectType,
				EffectName:       datum.EffectName,
				TrackingId:       datum.TrackingId,
				ClickTrackingUrl: datum.ClickTrackingUrl,
				ImpTrackingUrl:   datum.ImpTrackingUrl,
			})
			items = append(items, &TrackingItem{EffectType: datum.EffectType, EffectName: datum.EffectName, TrackingId: datum.TrackingId})
		} else {
			disabled = append(disabled, datum.TrackingId)
		}
	}
	err = model.NewTracking(vars.DBMysql).BatchInsert(trackings, disabled)
	if err != nil {
		return nil, 0, errors.New("写入失败：" + err.Error())
	}
	return items, response.Data.Total, nil
}
