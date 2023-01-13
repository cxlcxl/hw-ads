package serviceasset

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/library/curl"
	"errors"
)

type BaseAdsResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AdsAssetRequest struct {
	AdvertiserId string    `json:"advertiser_id"`
	Page         int       `json:"page"`
	PageSize     int64     `json:"page_size"`
	Filtering    Filtering `json:"filtering"`
}

type Filtering struct {
	AssetStatus string `json:"asset_status"`
}

type AdsAssetResponse struct {
	BaseAdsResp
	Data AdsAssetInfo `json:"data"`
}

type AdsAssetInfo struct {
	Total              int64        `json:"total"`
	CreativeAssetInfos []*AssetInfo `json:"creative_asset_infos"`
}

type AssetInfo struct {
	AssetStatus       string `json:"asset_status"`
	AssetId           int64  `json:"asset_id"`
	AssetName         string `json:"asset_name"`
	AssetType         string `json:"asset_type"`
	FileUrl           string `json:"file_url"`
	Width             int64  `json:"width"`
	Height            int64  `json:"height"`
	FileSize          int64  `json:"file_size"`
	FileFormat        string `json:"file_format"`
	FileHashSha256    string `json:"file_hash_sha256"`
	VideoPlayDuration int64  `json:"video_play_duration"`
}

func AssetPull(token string, d AdsAssetRequest, accountId int64) (int64, error) {
	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Asset.Query")).Get().JsonData(d)
	if err != nil {
		return 0, errors.New("华为接口生成错误：" + err.Error())
	}
	var response AdsAssetResponse
	if err = c.Request(&response, curl.Authorization(token)); err != nil {
		return 0, errors.New("华为接口调用错误：" + err.Error())
	}
	if response.Code != "200" {
		return 0, errors.New("华为接口拉取错误：" + response.Message)
	}
	assets := make([]*model.Asset, len(response.Data.CreativeAssetInfos))
	for i, datum := range response.Data.CreativeAssetInfos {
		assets[i] = &model.Asset{
			AccountId:         accountId,
			AssetId:           datum.AssetId,
			AssetName:         datum.AssetName,
			AssetType:         datum.AssetType,
			FileUrl:           datum.FileUrl,
			Width:             datum.Width,
			Height:            datum.Height,
			VideoPlayDuration: datum.VideoPlayDuration,
			FileSize:          datum.FileSize,
			FileFormat:        datum.FileFormat,
			FileHashSha256:    datum.FileHashSha256,
		}
	}
	err = model.NewAsset(vars.DBMysql).BatchInsert(assets)
	if err != nil {
		return 0, errors.New("数据写入失败：" + err.Error())
	}
	return response.Data.Total, nil
}
