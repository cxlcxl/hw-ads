package vars

const (
	ConfigKeyPrefix = "mobgi_ads_config_"
	LoginUserKey    = "mobgi_ads_login_user"

	UserStateValid = 1

	LoginTicketKey        = "mobgi:ads:app:ticket:"
	LoginTicketSessionKey = "LoginUserInfo"
	UserAllowModifySso    = "UserAllowModifySso"

	MaxPageSize       uint64 = 100
	SystemDefaultPass        = "a123456"

	DateTimeFormat = "2006-01-02 15:04:05"

	DateFormat = "2006-01-02"
	Env        = "dev"

	ApiModuleCountry         = "Country"
	ApiModuleCountryCollect  = "CountryCollect"
	ApiModuleAds             = "Ads"
	ApiModuleAdsCollect      = "AdsCollect"
	ApiModuleApp             = "App"
	ApiModuleCampaign        = "Campaign"
	ApiModuleDictionary      = "Dictionary"      // 定向字典数据
	ApiModuleRefreshToken    = "RefreshToken"    // Token 刷新
	ApiModuleTargeting       = "Targeting"       // 定向列表
	ApiModulePosition        = "Position"        // 版位
	ApiModulePositionPrice   = "PositionPrice"   // 版位底价
	ApiModulePositionElement = "PositionElement" // 版位元素

	// ScheduleTimeout 任务调度刷新间隔 /单位：s
	ScheduleTimeout = 300
)

const (
	ResponseCodeOk = iota
	ResponseCodeError
	ResponseCodeOvertime
	ResponseCodeDatabaseErr
	ResponseCodeValidFailed
	ResponseCodeUnauthorized
	ResponseCodeEmptyToken
	ResponseCodeTokenErr
	ResponseCodeTokenExpire
)

const (
	_ = iota
	AccountTypeMarket
	AccountTypeAds
)

const (
	_ = iota
	AppChannelGallery
	AppChannelGooglePlay
	AppChannelAppStore
	AppChannelOther
)

const (
	CommonStateVoid = iota
	CommonStateValid
)

// 调度截止规则：0 调度到当天；-1 停止调度此任务；> 0 为当前日期减{pause_rule}天
const (
	JobPauseToday = iota
	JobPauseADayAgo
	JobPauseTwoDayAgo
	JobPauseThreeDayAgo
	JobPauseFourDayAgo
	JobPauseFiveDayAgo
	JobPauseAWeekAgo
	JobPauseStop = -1
)

var (
	ResponseMsg = map[int]string{
		ResponseCodeOk:           "OK",
		ResponseCodeError:        "请求失败",
		ResponseCodeOvertime:     "请求超时",
		ResponseCodeDatabaseErr:  "数据库查询失败",
		ResponseCodeValidFailed:  "数据验证失败",
		ResponseCodeUnauthorized: "Unauthorized:权限不足",
		ResponseCodeEmptyToken:   "缺少 TOKEN",
		ResponseCodeTokenErr:     "TOKEN 错误",
		ResponseCodeTokenExpire:  "TOKEN 过期",
	}
	// AccountType 账号类型
	AccountType = map[int]string{
		AccountTypeMarket: "投放",
		AccountTypeAds:    "变现",
	}
	// CommonState 通用数据库状态字段
	CommonState = map[int]string{
		CommonStateVoid:  "停用",
		CommonStateValid: "正常",
	}
	// AppChannel 系统平台(渠道)
	AppChannel = map[int]string{
		AppChannelGallery:    "AppGallery",
		AppChannelGooglePlay: "GooglePlay",
		AppChannelAppStore:   "AppStore",
		AppChannelOther:      "Other",
	}
	JobPauseRule = map[int]string{
		JobPauseStop:        "停止任务调度",
		JobPauseToday:       "到当天",
		JobPauseADayAgo:     "一天前",
		JobPauseTwoDayAgo:   "两天前",
		JobPauseThreeDayAgo: "三天前",
		JobPauseFourDayAgo:  "四天前",
		JobPauseFiveDayAgo:  "五天前",
		JobPauseAWeekAgo:    "一周前",
	}
	ApiModules = map[string]string{
		ApiModuleCountry:         "[Country] 投放报表数据",
		ApiModuleCountryCollect:  "[CountryCollect] 投放报表数据整理",
		ApiModuleAds:             "[Ads] 变现报表数据",
		ApiModuleAdsCollect:      "[AdsCollect] 变现报表数据整理",
		ApiModuleApp:             "[App] 应用",
		ApiModuleCampaign:        "[Campaign] 投放任务数据",
		ApiModuleDictionary:      "[Dictionary] 定向字典数据",
		ApiModuleRefreshToken:    "[RefreshToken] Token 刷新",
		ApiModuleTargeting:       "[Targeting] 定向列表",
		ApiModulePosition:        "[Position] 版位",
		ApiModulePositionPrice:   "[PositionPrice] 版位底价",
		ApiModulePositionElement: "[PositionElement] 版位元素",
	}
)
