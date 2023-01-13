package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	servicemarketing "bs.mobgi.cc/app/service/marketing"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"strings"
)

type Marketing struct{}

func (h *Marketing) TrackingList(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VTrackingList)
	_tracks, err := model.NewTracking(vars.DBMysql).TrackingList(params.AppId)
	if err != nil {
		response.Fail(ctx, "转化跟踪获取失败："+err.Error())
		return
	}
	tracks := make([]*servicemarketing.TrackingItem, 0)
	if _tracks == nil || len(_tracks) == 0 {
		if tracks, err = servicemarketing.TrackingPull(params.AppId); err != nil {
			response.Fail(ctx, err.Error())
			return
		}
	} else {
		for _, tracking := range _tracks {
			tracks = append(tracks, &servicemarketing.TrackingItem{
				EffectName: tracking.EffectName,
				EffectType: tracking.EffectType,
				TrackingId: tracking.TrackingId,
			})
		}
	}

	response.Success(ctx, tracks)
}

func (h *Marketing) TrackingRefresh(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VTrackingRefresh)

	tracks, err := servicemarketing.TrackingPull(params.AppId)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	response.Success(ctx, tracks)
}

func (h *Marketing) DictQuery(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VDictQuery)
	keys := strings.Split(params.DictKey, ",")
	for _, key := range keys {
		if !utils.InArray(key, vars.TargetingDictionaryKeys) {
			response.Fail(ctx, "参数有误")
			return
		}
	}
	dictionaries, err := model.NewTargetingDict(vars.DBMysql).FindDictionaries(keys, nil)
	if err != nil {
		response.Fail(ctx, "参数有误："+err.Error())
		return
	}
	rs := make(map[string][]*servicemarketing.Dictionary)
	for _, dictionary := range dictionaries {
		if _, ok := rs[dictionary.DictKey]; ok {
			rs[dictionary.DictKey] = append(rs[dictionary.DictKey], &servicemarketing.Dictionary{
				Id:    dictionary.Id,
				Pid:   dictionary.Pid,
				Label: dictionary.Label,
				Value: dictionary.Value,
			})
		} else {
			item := make([]*servicemarketing.Dictionary, 1)
			item[0] = &servicemarketing.Dictionary{
				Id:    dictionary.Id,
				Pid:   dictionary.Pid,
				Label: dictionary.Label,
				Value: dictionary.Value,
			}
			rs[dictionary.DictKey] = item
		}
	}
	if _, ok := rs["app_category"]; ok {
		rs["app_category"] = servicemarketing.FormatAppCategory(rs["app_category"])
	}
	if _, ok := rs["app_interest"]; ok {
		rs["app_interest"] = servicemarketing.FormatAppInterest(rs["app_interest"])
	}
	if _, ok := rs["carrier"]; ok {
		countries, err := model.NewOverseasRegion(vars.DBMysql).GetCountries()
		if err != nil {
			response.Fail(ctx, "国家数据为空")
			return
		}
		rs["carrier"] = servicemarketing.FormatCarrier(rs["carrier"], countries)
	}
	if _, ok := rs["not_pre_define_audience"]; ok {
		rs["not_pre_define_audience"] = servicemarketing.FormatAudience(rs["not_pre_define_audience"])
	}
	if _, ok := rs["pre_define_audience"]; ok {
		rs["pre_define_audience"] = servicemarketing.FormatAudience(rs["pre_define_audience"])
	}
	if _, ok := rs["media_app_category"]; ok {
		rs["media_app_category"] = servicemarketing.FormatMediaAppCategory(rs["media_app_category"])
	}
	response.Success(ctx, rs)
}
