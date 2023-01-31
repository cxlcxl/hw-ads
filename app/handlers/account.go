package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	serviceaccount "bs.mobgi.cc/app/service/account"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Account struct{}

func (h *Account) AccountList(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountList)
	offset := utils.GetPages(params.Page, params.PageSize)
	acts, total, err := model.NewAct(vars.DBMysql).AccountList(params.AccountType, params.State, params.AccountName, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"total": total,
		"list":  acts,
		"state": vars.CommonState,
		"types": vars.AccountType,
	})
}

func (h *Account) AccountParents(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountParents)
	acts, err := model.NewAct(vars.DBMysql).RemoteAccounts(params.AccountName, 1)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	ids := make([]int64, len(acts))
	for i, act := range acts {
		ids[i] = act.Id
	}
	if params.ParentId > 0 && !utils.InArray(params.ParentId, ids) {
		if act, err := model.NewAct(vars.DBMysql).FindAccountById(params.ParentId); err == nil {
			acts = append(acts, act)
		}
	}
	response.Success(ctx, acts)
}

func (h *Account) AccountDefault(ctx *gin.Context) {
	acts, err := model.NewAct(vars.DBMysql).RemoteAccounts("", 0)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, acts)
}

func (h *Account) AllAccounts(ctx *gin.Context) {
	acts, err := model.NewAct(vars.DBMysql).AllAccounts()
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, acts)
}

func (h *Account) AccountSearch(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountSearch)
	if params.AccountName == "" {
		response.Success(ctx, []interface{}{})
		return
	}
	acts, err := model.NewAct(vars.DBMysql).RemoteAccounts(params.AccountName, 0)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	response.Success(ctx, acts)
}

func (h *Account) AccountInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	act, err := model.NewAct(vars.DBMysql).FindAccountById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, act)
}

func (h *Account) AccountAuth(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("id"), 0, 64)
	if id <= 0 {
		response.Fail(ctx, "请求失败")
		return
	}
	info, err := model.NewAct(vars.DBMysql).FindAccountById(id)
	if err != nil {
		response.Fail(ctx, "请求有误")
		return
	}
	clientId, secret := info.ClientId, info.Secret
	if clientId == "" || secret == "" {
		if info.ParentId == 0 {
			response.Fail(ctx, "请先完整填写 ClientId 与 Secret")
			return
		} else {
			parent, err := model.NewAct(vars.DBMysql).FindAccountById(info.ParentId)
			if err != nil {
				response.Fail(ctx, "上级信息查询有误")
				return
			}
			if parent.ClientId == "" || parent.Secret == "" {
				response.Fail(ctx, "上级 ClientId 与 Secret 信息也未填写")
				return
			}
			clientId, secret = parent.ClientId, parent.Secret
		}
	}
	url, err := serviceaccount.AuthorizeCodeUrl(info.Id, clientId, secret)
	if err != nil {
		response.Fail(ctx, "请求失败:"+err.Error())
		return
	}
	response.Success(ctx, url)
}

func (h *Account) RefreshToken(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	token, err := model.NewToken(vars.DBMysql).FindByAccountId(id)
	if err != nil {
		response.Fail(ctx, "token 查询失败："+err.Error())
		return
	}
	info, err := model.NewAct(vars.DBMysql).FindAccountById(id)
	if err != nil {
		response.Fail(ctx, "账户查询失败："+err.Error())
		return
	}
	clientId, secret := info.ClientId, info.Secret
	if clientId == "" || secret == "" {
		if info.ParentId == 0 {
			response.Fail(ctx, "请先完整填写 ClientId 与 Secret")
			return
		} else {
			parent, err := model.NewAct(vars.DBMysql).FindAccountById(info.ParentId)
			if err != nil {
				response.Fail(ctx, "父账户查询失败："+err.Error())
				return
			}
			if parent.ClientId == "" || parent.Secret == "" {
				response.Fail(ctx, "检查上级 ClientId 信息有误，请检查是否有完整填写")
				return
			}
			clientId, secret = parent.ClientId, parent.Secret
		}
	}
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
		"client_id":     clientId,
		"client_secret": secret,
	}
	var at vars.AdsAccessTokenResponse
	refreshUrl := vars.YmlConfig.GetString("MarketingApis.Refresh")
	if err = curl.New(refreshUrl).Post().QueryData(data).Request(&at, curl.FormHeader()); err != nil {
		response.Fail(ctx, "请求构建失败："+err.Error())
		return
	}
	if at.Error != 0 {
		response.Fail(ctx, "华为接口调用失败："+at.ErrorDescription)
		return
	}
	expire := time.Now().Unix() + at.ExpiresIn - 20
	err = model.NewToken(vars.DBMysql).TokenUpdate(map[string]interface{}{
		"access_token": at.AccessToken,
		"expired_at":   time.Unix(expire, 0),
		"updated_at":   time.Now(),
		"token_type":   at.TokenType,
	}, token.Id)
	if err != nil {
		response.Fail(ctx, "更细腻失败："+err.Error())
		return
	}

	response.Success(ctx, &vars.AccountTokenInfo{
		AccountId:    id,
		AccessToken:  at.AccessToken,
		RefreshToken: at.RefreshToken,
		ExpiredAt:    expire,
		TokenType:    at.TokenType,
	})
}

func (h *Account) AccountUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountUpdate)
	d := map[string]interface{}{
		"account_name":  params.AccountName,
		"account_type":  params.AccountType,
		"advertiser_id": params.AdvertiserId,
		"developer_id":  params.DeveloperId,
		"state":         params.State,
		"client_id":     params.ClientId,
		"secret":        params.Secret,
		"parent_id":     params.ParentId,
		"updated_at":    time.Now(),
	}
	if err := model.NewAct(vars.DBMysql).AccountUpdate(d, params.Id); err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Account) AccountCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountCreate)
	d := &model.Account{
		ParentId:     params.ParentId,
		AdvertiserId: params.AdvertiserId,
		DeveloperId:  params.DeveloperId,
		AccountType:  params.AccountType,
		State:        1,
		AccountName:  params.AccountName,
		ClientId:     params.ClientId,
		Secret:       params.Secret,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := model.NewAct(vars.DBMysql).AccountCreate(d); err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, nil)
}
