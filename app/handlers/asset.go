package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	serviceaccount "bs.mobgi.cc/app/service/account"
	serviceasset "bs.mobgi.cc/app/service/asset"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Asset struct{}

func (h *Asset) AssetList(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAssetList)
	offset := utils.GetPages(params.Page, params.PageSize)
	list, total, err := model.NewAsset(vars.DBMysql).AssetList(params.AppId, params.AssetType, params.Width, params.Height, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, gin.H{"asset_type": vars.AssetType, "total": total, "list": list})
}

func (h *Asset) AssetDimension(ctx *gin.Context) {
	var (
		pictureDimension []string
		videoDimension   []string
	)
	for _, dimension := range vars.PictureDimensions {
		pictureDimension = append(pictureDimension, fmt.Sprintf("%d*%d", dimension.Width, dimension.Height))
	}
	for _, dimension := range vars.VideoDimensions {
		videoDimension = append(videoDimension, fmt.Sprintf("%d*%d", dimension.Width, dimension.Height))
	}
	response.Success(ctx, gin.H{
		vars.AssetTypePicture: pictureDimension,
		vars.AssetTypeVideo:   videoDimension,
	})
}

func (h *Asset) AssetSync(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAssetSync)
	app, err := model.NewApp(vars.DBMysql).FindAppByAppId(params.AppId)
	if err != nil {
		response.Fail(ctx, "应用数据异常："+err.Error())
		return
	}
	token, err := serviceaccount.GetToken(app.AccountId)
	if err != nil {
		response.Fail(ctx, "Token 信息获取失败："+err.Error())
		return
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		response.Fail(ctx, "Token 已过期，请先到账户列表页刷新 Token")
		return
	}
	page := 1
	d := serviceasset.AdsAssetRequest{
		Page:      page,
		PageSize:  50,
		Filtering: serviceasset.Filtering{AssetStatus: "CREATIVE_ASSET_ENABLE"},
	}
	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	total, err := serviceasset.AssetPull(t, d, app.AccountId)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	if total > 50 {
		pages := utils.CeilPages(total, 50)
		for i := 2; i <= int(pages); i++ {
			d.Page = i
			if _, err = serviceasset.AssetPull(t, d, app.AccountId); err != nil {
				response.Fail(ctx, err.Error())
				return
			}
		}
	}
	response.Success(ctx, nil)
}
