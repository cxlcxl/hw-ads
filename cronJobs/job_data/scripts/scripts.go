package scripts

import (
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/jobs/ads"
	"bs.mobgi.cc/cronJobs/jobs/app"
	"bs.mobgi.cc/cronJobs/jobs/campaign"
	"bs.mobgi.cc/cronJobs/jobs/country"
	"bs.mobgi.cc/cronJobs/jobs/dictionary"
	"bs.mobgi.cc/cronJobs/jobs/position"
	"bs.mobgi.cc/cronJobs/jobs/refreshToken"
	"bs.mobgi.cc/cronJobs/jobs/targeting"
	"time"
)

type Job func()
type ManualJob func(day time.Time, pauseRule int64)

var (
	ScheduleJobs = map[string]Job{
		vars.ApiModuleCountry:         country.Country,
		vars.ApiModuleCampaign:        campaign.Campaign,
		vars.ApiModuleDictionary:      dictionary.Dictionary,
		vars.ApiModuleRefreshToken:    refreshToken.RefreshToken,
		vars.ApiModuleTargeting:       targeting.Targeting,
		vars.ApiModulePosition:        position.Position,
		vars.ApiModulePositionPrice:   position.Price,
		vars.ApiModulePositionElement: position.Element,
		vars.ApiModuleAds:             ads.ReportAds,
		vars.ApiModuleApp:             app.ReportApp,
	}

	ManualScheduleJobs = map[string]ManualJob{
		vars.ApiModuleCampaign:        campaign.CampaignManual,
		vars.ApiModuleRefreshToken:    refreshToken.RefreshTokenManual,
		vars.ApiModulePositionPrice:   position.PriceManual,
		vars.ApiModulePositionElement: position.ElementManual,
		vars.ApiModuleCountry:         country.CountryManual,
		vars.ApiModuleAds:             ads.ReportAdsManual,
		vars.ApiModuleApp:             app.ReportAppManual,
	}
)
