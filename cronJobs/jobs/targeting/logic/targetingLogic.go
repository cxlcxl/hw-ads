package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/statements"
	"bs.mobgi.cc/cronJobs/jobs"
	"bs.mobgi.cc/library/curl"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type TargetingLogic struct {
	tokenChan chan *jobs.QueryParam
	wg        sync.WaitGroup
	pageSize  int64
}

func NewTargetingLogic() *TargetingLogic {
	return &TargetingLogic{
		tokenChan: make(chan *jobs.QueryParam),
		wg:        sync.WaitGroup{},
		pageSize:  50,
	}
}

func (l *TargetingLogic) TargetingQuery() (err error) {
	go jobs.GetTokens(l.tokenChan)

	for token := range l.tokenChan {
		account, err := model.NewAct(vars.DBMysql).FindAccountById(token.AccountId) // 只有几条数据
		if err != nil {
			fmt.Println("定向包调度失败：", err)
		}
		if err = l.query(token, account.AdvertiserId, 1); err != nil {
			fmt.Println("定向包调度失败：", err)
		}
	}

	return
}

func (l *TargetingLogic) query(param *jobs.QueryParam, advertiserId string, page int64) (err error) {
	data := statements.TargetingRequest{
		RequestPage:  statements.RequestPage{Page: page, PageSize: l.pageSize},
		AdvertiserId: advertiserId,
	}
	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Tools.Targeting")).Get().JsonData(data)
	if err != nil {
		return err
	}
	var response statements.TargetingResponse
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		return err
	}
	if response.Code != "200" {
		return errors.New("接口返回错误：" + response.Message)
	}
	targetings := formatTargeting(response.Data.Targetings, param.AccountId, advertiserId)
	if err = model.NewTargeting(vars.DBMysql).BatchInsert(targetings); err != nil {
		fmt.Println("写入数据错误：", err)
	}
	offset := page * l.pageSize
	if response.Data.Total > offset {
		page++
		if err = l.query(param, advertiserId, page); err != nil {
			fmt.Println("定向包分页调度失败：", page, err)
		}
	}
	return err
}

func formatTargeting(src []*statements.Targeting, accountId int64, advertiserId string) (targetings []*model.Targeting) {
	for _, targeting := range src {
		locationType, include, exclude := "", "", ""
		if !isEmpty(targeting.CurrentCustomLocation) {
			locationType = "current"
			include = targetingValue(targeting.CurrentCustomLocation)
		}
		if !isEmpty(targeting.NotCurrentCustomLocation) {
			locationType = "current"
			exclude = targetingValue(targeting.NotCurrentCustomLocation)
		}
		if !isEmpty(targeting.ResidenceCustomLocation) {
			locationType = "residence"
			include = targetingValue(targeting.ResidenceCustomLocation)
		}
		if !isEmpty(targeting.NotResidenceCustomLocation) {
			locationType = "residence"
			exclude = targetingValue(targeting.NotResidenceCustomLocation)
		}
		carriers := ""
		if !isEmpty(targeting.Carrier) {
			carriers = targetingValue(targeting.Carrier)
		}
		language := ""
		if !isEmpty(targeting.Language) {
			language = targetingValue(targeting.Language)
		}
		age := ""
		if !isEmpty(targeting.Age) {
			age = targetingValue(targeting.Age)
		}
		gender := ""
		if !isEmpty(targeting.Gender) {
			gender = targetingValue(targeting.Gender)
		}
		appCategory, appCategories := "", ""
		if !isEmpty(targeting.AppCategoryActive) {
			appCategory = "1"
			appCategories = targetingValue(targeting.AppCategoryActive)
		}
		if !isEmpty(targeting.AppCategoryInstalled) {
			appCategory = "2"
			appCategories = targetingValue(targeting.AppCategoryInstalled)
		}
		if !isEmpty(targeting.NotAppCategoryInstall) {
			appCategory = "3"
			appCategories = targetingValue(targeting.NotAppCategoryInstall)
		}
		installApp := ""
		if !isEmpty(targeting.InstalledApps) {
			installApp = "1"
		}
		if !isEmpty(targeting.NotInstalledApps) {
			installApp = "0"
		}
		appInterest, appInterests := "", ""
		if !isEmpty(targeting.UnLimitAppInterest) {
			appInterest = "1"
			appInterests = targetingValue(targeting.UnLimitAppInterest)
		}
		if !isEmpty(targeting.NormalAppInterest) {
			appInterest = "2"
			appInterests = targetingValue(targeting.NormalAppInterest)
		}
		if !isEmpty(targeting.HighAppInterest) {
			appInterest = "3"
			appInterests = targetingValue(targeting.HighAppInterest)
		}
		series := ""
		if !isEmpty(targeting.SeriesType) {
			series = targetingValue(targeting.SeriesType)
		}
		network := ""
		if !isEmpty(targeting.NetworkType) {
			network = targetingValue(targeting.NetworkType)
		}
		audiences, notAudiences := "", ""
		if !isEmpty(targeting.Audience) {
			audiences = targetingValue(targeting.Audience)
		}
		if !isEmpty(targeting.NotAudience) {
			notAudiences = targetingValue(targeting.NotAudience)
		}
		mediaAppCategory := ""
		if !isEmpty(targeting.AppCategoryOfMedia) {
			mediaAppCategory = targetingValue(targeting.AppCategoryOfMedia)
		}
		targetings = append(targetings, &model.Targeting{
			AccountId:          accountId,
			AdvertiserId:       advertiserId,
			TargetingId:        targeting.TargetingId,
			TargetingName:      targeting.TargetingName,
			TargetingType:      targeting.TargetingType,
			LocationType:       locationType,
			IncludeLocation:    include,
			ExcludeLocation:    exclude,
			Carriers:           carriers,
			Language:           language,
			Age:                age,
			Gender:             gender,
			AppCategory:        appCategory,
			AppCategories:      appCategories,
			InstalledApps:      installApp,
			AppInterest:        appInterest,
			AppInterests:       appInterests,
			Series:             series,
			NetworkType:        network,
			NotAudiences:       notAudiences,
			Audiences:          audiences,
			AppCategoryOfMedia: mediaAppCategory,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		})
	}
	return
}

func isEmpty(v statements.TargetingValue) bool {
	if len(v.Value) == 0 {
		return true
	}
	if v.Value[0] == "" {
		return true
	}
	return false
}

func targetingValue(v statements.TargetingValue) string {
	return strings.Join(v.Value, vars.TargetingDatabaseSeq)
}
