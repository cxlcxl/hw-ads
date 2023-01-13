package servicemarketing

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"errors"
	"strings"
	"time"
)

type TargetingListItem struct {
	TargetingId int64 `json:"targeting_id"`
	TargetingCreate
}
type TargetingCreate struct {
	CampaignId         string     `json:"campaign_id"`
	TargetingType      string     `json:"targeting_type"`
	TargetingName      string     `json:"targeting_name"`
	Gender             string     `json:"gender"`
	Age                []string   `json:"age"`
	NetworkType        []string   `json:"network_type"`
	Location           string     `json:"location"`
	LocationType       string     `json:"location_type"`
	IncludeLocation    []string   `json:"include_location"`
	ExcludeLocation    []string   `json:"exclude_location"`
	Carrier            string     `json:"carrier"`
	Carriers           [][]string `json:"carriers"`
	AppCategory        string     `json:"app_category"`
	AppCategories      []string   `json:"app_categories"`
	AppInterest        string     `json:"app_interest"`
	AppInterests       [][]string `json:"app_interests"`
	Audience           string     `json:"audience"`
	Audiences          []string   `json:"audiences"`
	NotAudience        []string   `json:"not_audience"`
	SeriesType         string     `json:"series_type"`
	Series             []string   `json:"series"`
	MediaAppCategory   string     `json:"media_app_category"`
	AppCategoryOfMedia [][]string `json:"app_category_of_media"`
	LanguageCheck      string     `json:"language_check"`
	Language           []string   `json:"language"`
	InstalledApps      string     `json:"installed_apps"`
}

func UnFormatTargetData(src []*model.Targeting) (targetings []*TargetingListItem, err error) {
	carriers := make([]string, 0)
	appInterests := make([]string, 0)
	mediaAppCategories := make([]string, 0)
	targetings = make([]*TargetingListItem, len(src))
	for i, targeting := range src {
		var (
			appCategories = make([]string, 0)
			include       = make([]string, 0)
			exclude       = make([]string, 0)
			audiences     = make([]string, 0)
			notAudiences  = make([]string, 0)
			series        = make([]string, 0)
			language      = make([]string, 0)
		)
		location := ""
		if targeting.IncludeLocation != "" || targeting.ExcludeLocation != "" {
			location = "1"
			include = targetingSplit(targeting.IncludeLocation)
			exclude = targetingSplit(targeting.ExcludeLocation)
		}
		audience := ""
		if targeting.Audiences != "" || targeting.NotAudiences != "" {
			audience = "1"
			audiences = targetingSplit(targeting.Audiences)
			notAudiences = targetingSplit(targeting.NotAudiences)
		}
		carrier := ""
		if targeting.Carriers != "" {
			carrier = "1"
			carriers = append(carriers, targetingSplit(targeting.Carriers)...)
		}
		seriesType := ""
		if targeting.Series != "" {
			seriesType = "1"
			series = targetingSplit(targeting.Series)
		}
		if targeting.AppCategory != "" {
			appCategories = targetingSplit(targeting.AppCategories)
		}
		if targeting.AppInterest != "" {
			appInterestsTmp := targetingSplit(targeting.AppInterests)
			appInterests = append(appInterests, appInterestsTmp...)
		}
		mediaAppCategory := ""
		if targeting.AppCategoryOfMedia != "" {
			mediaAppCategory = "1"
			mediaAppCategories = append(mediaAppCategories, targetingSplit(targeting.AppCategoryOfMedia)...)
		}
		languageCheck := ""
		if targeting.Language != "" {
			languageCheck = "1"
			language = targetingSplit(targeting.Language)
		}
		targetings[i] = &TargetingListItem{
			TargetingId: targeting.TargetingId,
			TargetingCreate: TargetingCreate{
				TargetingType:    targeting.TargetingType,
				TargetingName:    targeting.TargetingName,
				Gender:           targeting.Gender,
				Age:              strings.Split(targeting.Age, vars.TargetingDatabaseSeq),
				NetworkType:      strings.Split(targeting.NetworkType, vars.TargetingDatabaseSeq),
				Location:         location,
				LocationType:     targeting.LocationType,
				IncludeLocation:  include,
				ExcludeLocation:  exclude,
				Carrier:          carrier,
				AppCategory:      targeting.AppCategory,
				AppCategories:    appCategories,
				AppInterest:      targeting.AppInterest,
				Audience:         audience,
				Audiences:        audiences,
				NotAudience:      notAudiences,
				SeriesType:       seriesType,
				Series:           series,
				MediaAppCategory: mediaAppCategory,
				LanguageCheck:    languageCheck,
				Language:         language,
				InstalledApps:    targeting.InstalledApps,
			},
		}
	}
	// 运营商
	carrierMap := make(map[string]string)
	carrierParents := make(map[string]string)
	if len(carriers) > 0 {
		dictCarriers, err := model.NewTargetingDict(vars.DBMysql).FindDictionaries([]string{"carrier"}, carriers)
		if err != nil {
			return nil, errors.New("运营商解析失败：" + err.Error())
		}
		countries, err := model.NewOverseasRegion(vars.DBMysql).GetCountries()
		if err != nil {
			return nil, errors.New("国家数据获取失败：" + err.Error())
		}
		for _, dict := range dictCarriers {
			carrierMap[dict.Value] = dict.Pid
		}
		for _, country := range countries {
			carrierParents[country.CCode] = country.CId
		}
	}

	// 媒体类型
	mediaCategoryParents := make(map[string]string)
	if len(mediaAppCategories) > 0 {
		mediaCategories, err := model.NewTargetingDict(vars.DBMysql).FindDictionaries([]string{"media_app_category"}, mediaAppCategories)
		if err != nil {
			return nil, errors.New("媒体类型解析失败：" + err.Error())
		}
		for _, dictionary := range mediaCategories {
			mediaCategoryParents[dictionary.Value] = dictionary.Pid
		}
	}

	// App 兴趣
	appInterestsParents := make(map[string]string)
	interestParentMap := make(map[string]string)
	if len(appInterests) > 0 {
		dictAppInterests, err := model.NewTargetingDict(vars.DBMysql).FindDictionaries([]string{"app_interest"}, appInterests)
		if err != nil {
			return nil, errors.New("App 兴趣解析失败：" + err.Error())
		}
		for _, dictionary := range dictAppInterests {
			appInterestsParents[dictionary.Value] = dictionary.Pid
		}
		for s, s2 := range vars.AppInterest {
			interestParentMap[s2] = s
		}
	}
	// 组合数据
	for i, targeting := range src {
		carriers = targetingSplit(targeting.Carriers)
		mediaAppCategories = targetingSplit(targeting.AppCategoryOfMedia)
		appInterests = targetingSplit(targeting.AppInterests)
		targetingCarriers := make([][]string, 0)
		targetingMediaCategories := make([][]string, 0)
		targetingAppInterests := make([][]string, 0)

		for _, carrier := range carriers {
			if cId, ok := carrierParents[carrierMap[carrier]]; !ok {
				return nil, errors.New("字典与定向包数据匹配异常：运营商")
			} else {
				targetingCarriers = append(targetingCarriers, []string{cId, carrier})
			}
		}
		for _, category := range mediaAppCategories {
			if pid, ok := mediaCategoryParents[category]; !ok {
				return nil, errors.New("字典与定向包数据匹配异常：媒体类型")
			} else {
				targetingMediaCategories = append(targetingMediaCategories, []string{pid, category})
			}
		}
		for _, appInterest := range appInterests {
			if pid, ok := appInterestsParents[appInterest]; !ok {
				return nil, errors.New("字典与定向包数据匹配异常：App 兴趣")
			} else {
				targetingAppInterests = append(targetingAppInterests, []string{interestParentMap[pid], appInterest})
			}
		}

		targetings[i].Carriers = targetingCarriers
		targetings[i].AppCategoryOfMedia = targetingMediaCategories
		targetings[i].AppInterests = targetingAppInterests
	}
	return
}

func targetingSplit(s string) []string {
	if s == "" {
		return []string{}
	} else {
		return strings.Split(s, vars.TargetingDatabaseSeq)
	}
}

type TargetingValue struct {
	Value []string `json:"value"`
}

type TargetingCreateResp struct {
	BaseAdsResp
	Data struct {
		TargetingId int64 `json:"targeting_id"`
	} `json:"data"`
}

func FormatAdsData(req *v_data.VTargetingCreate, advertiserId string) (rs map[string]interface{}, err error) {
	rs = make(map[string]interface{})
	rs["targeting_name"] = req.TargetingName
	rs["advertiser_id"] = advertiserId
	rs["targeting_type"] = vars.TargetingTypeApp
	// 地域
	if req.Location == "1" {
		if req.LocationType == "current" {
			if len(req.IncludeLocation) > 0 {
				rs["current_custom_location_struct"] = TargetingValue{Value: req.IncludeLocation}
			}
			if len(req.ExcludeLocation) > 0 {
				rs["not_current_custom_location_struct"] = TargetingValue{Value: req.ExcludeLocation}
			}
		} else if req.LocationType == "residence" {
			if len(req.IncludeLocation) > 0 {
				rs["residence_custom_location_struct"] = TargetingValue{Value: req.IncludeLocation}
			}
			if len(req.ExcludeLocation) > 0 {
				rs["not_residence_custom_location_struct"] = TargetingValue{Value: req.ExcludeLocation}
			}
		} else {
			return rs, errors.New("未限制地域类型，或地域选择的类型错误")
		}
	}
	// 性别
	if req.Gender != "" {
		rs["gender_struct"] = TargetingValue{Value: []string{req.Gender}}
	}
	// 年龄
	if len(req.Age) > 0 {
		if !utils.InArray("", req.Age) {
			rs["age_struct"] = TargetingValue{Value: req.Age}
		}
	}
	// App 安装
	if req.InstalledApps == "1" {
		rs["installed_apps_struct"] = TargetingValue{Value: []string{"true"}}
	} else {
		rs["not_installed_apps_struct"] = TargetingValue{Value: []string{"true"}}
	}
	// App 行为
	if req.AppCategory != "" {
		switch req.AppCategory {
		case "1":
			rs["app_category_active_struct"] = TargetingValue{Value: req.AppCategories}
			break
		case "2":
			rs["app_category_install_struct"] = TargetingValue{Value: req.AppCategories}
			break
		case "3":
			rs["not_app_category_install_struct"] = TargetingValue{Value: req.AppCategories}
			break
		default:
			return rs, errors.New("App 行为选择错误")
		}
	}
	// App 兴趣
	if req.AppInterest != "" {
		var appInterests []string
		for _, interest := range req.AppInterests {
			appInterests = append(appInterests, interest[len(interest)-1])
		}
		switch req.AppInterest {
		case "1":
			rs["unlimit_app_interest_struct"] = TargetingValue{Value: appInterests}
			break
		case "2":
			rs["normal_app_interest_struct"] = TargetingValue{Value: appInterests}
			break
		case "3":
			rs["high_app_interest_struct"] = TargetingValue{Value: appInterests}
			break
		default:
			return rs, errors.New("App 兴趣选择错误")
		}
	}
	// 设备
	if req.SeriesType == "1" {
		rs["series_type_struct"] = TargetingValue{Value: req.Series}
	}
	// 联网方式
	if len(req.NetworkType) > 0 {
		if !utils.InArray("", req.NetworkType) {
			rs["network_type_struct"] = TargetingValue{Value: req.NetworkType}
		}
	}
	// 自定义人群
	if req.Audience == "1" {
		rs["audience_struct"] = TargetingValue{Value: req.Audiences}
		rs["not_audience_struct"] = TargetingValue{Value: req.NotAudience}
	}
	// 媒体类型
	if req.MediaAppCategory == "1" {
		var categories []string
		for _, appCategoryOfMedia := range req.AppCategoryOfMedia {
			categories = append(categories, appCategoryOfMedia[len(appCategoryOfMedia)-1])
		}
		rs["app_category_of_media_struct"] = TargetingValue{Value: categories}
	}
	// 语言
	if req.LanguageCheck == "1" {
		rs["language_struct"] = TargetingValue{Value: req.Language}
	}
	// 运营商
	if req.Carrier == "1" {
		var carriers []string
		for _, carrier := range req.Carriers {
			carriers = append(carriers, carrier[len(carrier)-1])
		}
		rs["carrier_struct"] = TargetingValue{Value: carriers}
	}

	return
}

func FormatRDBData(req *v_data.VTargetingCreate, advertiserId string, targetingId, accountId int64) (targeting *model.Targeting) {
	// 地域
	include, exclude, appInterests, mediaCategories, categories, series := "", "", "", "", "", ""
	if req.Location == "1" {
		include = strings.Join(req.IncludeLocation, vars.TargetingDatabaseSeq)
		exclude = strings.Join(req.ExcludeLocation, vars.TargetingDatabaseSeq)
	}
	// App 兴趣
	if req.AppInterest != "" {
		var tmp []string
		for _, interest := range req.AppInterests {
			tmp = append(tmp, interest[len(interest)-1])
		}
		appInterests = strings.Join(tmp, vars.TargetingDatabaseSeq)
	}
	// 媒体类型
	if req.MediaAppCategory == "1" {
		var tmp []string
		for _, appCategoryOfMedia := range req.AppCategoryOfMedia {
			tmp = append(tmp, appCategoryOfMedia[len(appCategoryOfMedia)-1])
		}
		mediaCategories = strings.Join(tmp, vars.TargetingDatabaseSeq)
	}
	// 运营商
	carriers := ""
	if req.Carrier == "1" {
		var tmp []string
		for _, carrier := range req.Carriers {
			tmp = append(tmp, carrier[len(carrier)-1])
		}
		carriers = strings.Join(tmp, vars.TargetingDatabaseSeq)
	}
	// 语言
	language := ""
	if req.LanguageCheck == "1" {
		language = strings.Join(req.Language, vars.TargetingDatabaseSeq)
	}
	// App 行为
	if req.AppCategory != "" {
		categories = strings.Join(req.AppCategories, vars.TargetingDatabaseSeq)
	}
	// 设备
	if req.SeriesType == "1" {
		series = strings.Join(req.Series, vars.TargetingDatabaseSeq)
	}
	// 自定义人群
	audiences, notAudiences := "", ""
	if req.Audience == "1" {
		audiences = strings.Join(req.Audiences, vars.TargetingDatabaseSeq)
		notAudiences = strings.Join(req.NotAudience, vars.TargetingDatabaseSeq)
	}
	targeting = &model.Targeting{
		AccountId:          accountId,
		AdvertiserId:       advertiserId,
		TargetingId:        targetingId,
		TargetingName:      req.TargetingName,
		TargetingType:      vars.TargetingTypeApp,
		LocationType:       req.LocationType,
		IncludeLocation:    include,
		ExcludeLocation:    exclude,
		Carriers:           carriers,
		Language:           language,
		Age:                strings.Join(req.Age, vars.TargetingDatabaseSeq),
		Gender:             req.Gender,
		AppCategory:        req.AppCategory,
		AppCategories:      categories,
		InstalledApps:      req.InstalledApps,
		AppInterest:        req.AppInterest,
		AppInterests:       appInterests,
		Series:             series,
		NetworkType:        strings.Join(req.NetworkType, vars.TargetingDatabaseSeq),
		NotAudiences:       notAudiences,
		Audiences:          audiences,
		AppCategoryOfMedia: mediaCategories,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	return
}
