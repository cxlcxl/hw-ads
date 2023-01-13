package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type App struct{}

func (h *App) AppList(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAppList)
	offset := utils.GetPages(params.Page, params.PageSize)
	apps, total, err := model.NewApp(vars.DBMysql).AppList(params.AppId, params.AppName, params.Channel, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询失败："+err.Error())
		return
	}
	actIds := make([]int64, 0)
	for _, app := range apps {
		actIds = append(actIds, app.AccountId)
	}
	if acts, err := model.NewAct(vars.DBMysql).GetAccountsByIds(actIds); err == nil {
		_act := make(map[int64]string)
		for _, act := range acts {
			_act[act.Id] = act.AccountName
		}
		for i, app := range apps {
			if v, ok := _act[app.AccountId]; ok {
				apps[i].AccountName = v
			}
		}
	}
	response.Success(ctx, gin.H{
		"total":       total,
		"list":        apps,
		"app_channel": vars.AppChannel,
	})
}

func (h *App) AppCampaignList(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAppCampaignList)
	offset := utils.GetPages(params.Page, params.PageSize)
	apps, total, err := model.NewApp(vars.DBMysql).AppCampaignList(params.AppName, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询失败："+err.Error())
		return
	}
	response.Success(ctx, gin.H{
		"total":      total,
		"list":       apps,
		"total_page": utils.CeilPages(total, params.PageSize),
	})
}

func (h *App) AppUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAppUpdate)
	d := map[string]interface{}{
		"app_name":   params.AppName,
		"channel":    params.Channel,
		"pkg_name":   params.PkgName,
		"tags":       params.Tags,
		"account_id": params.AccountId,
		"updated_at": time.Now(),
	}
	err := model.NewApp(vars.DBMysql).UpdateApp(d, params.Id)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (h *App) AppCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAppCreate)
	app := &model.App{
		AccountId: params.AccountId,
		AppId:     params.AppId,
		AppName:   params.AppName,
		PkgName:   params.PkgName,
		Channel:   params.Channel,
		Tags:      params.Tags,
		Timestamp: model.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err := model.NewApp(vars.DBMysql).CreateApp(app)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (h *App) AppInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误："+err.Error())
		return
	}
	app, err := model.NewApp(vars.DBMysql).FindAppById(id)
	if err != nil {
		response.Fail(ctx, "查询失败："+err.Error())
		return
	}
	response.Success(ctx, app)
}

func (h *App) AppPull(ctx *gin.Context, p interface{}) {
	response.Success(ctx, nil)
}
