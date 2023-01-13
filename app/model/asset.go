package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Asset struct {
	connectDb

	Id                int64  `json:"id"`
	AccountId         int64  `json:"account_id"`    // 账户 ID
	AdvertiserId      string `json:"advertiser_id"` // 广告主账户 ID
	AppId             string `json:"app_id"`        // 第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位
	AssetId           int64  `json:"asset_id"`
	AssetName         string `json:"asset_name"`
	AssetType         string `json:"asset_type"`
	FileUrl           string `json:"file_url"`
	Width             int64  `json:"width"`
	Height            int64  `json:"height"`
	VideoPlayDuration int64  `json:"video_play_duration"`
	FileSize          int64  `json:"file_size"`
	FileFormat        string `json:"file_format"`
	FileHashSha256    string `json:"file_hash_sha256"`
}

func (m *Asset) TableName() string {
	return "assets"
}

func NewAsset(db *gorm.DB) *Asset {
	return &Asset{connectDb: connectDb{DB: db}}
}

func (m *Asset) AssetList(appId, assetType string, w, h int64, offset, limit int64) (assets []*Asset, total int64, err error) {
	query := m.Table(m.TableName()).Where("app_id = ?", appId)
	if len(assetType) > 0 {
		query = query.Where("asset_type = ?", assetType)
	}
	if w > 0 {
		query = query.Where("width = ?", w)
	}
	if h > 0 {
		query = query.Where("height = ?", h)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = query.Offset(int(offset)).Limit(int(limit)).Order("width asc,height asc,asset_type asc").Find(&assets).Error
	}
	return
}

func (m *Asset) BatchInsert(assets []*Asset) (err error) {
	assetsRowsExpectAutoSet := "account_id,advertiser_id,app_id,asset_id,asset_name,asset_type,file_url,width,height,video_play_duration,file_size,file_format,file_hash_sha256"
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.TableName(), assetsRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	for i := 0; i < len(assets); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		values = append(values,
			assets[i].AccountId, assets[i].AdvertiserId, assets[i].AppId, assets[i].AssetId, assets[i].AssetName,
			assets[i].AssetType, assets[i].FileUrl, assets[i].Width, assets[i].Height, assets[i].VideoPlayDuration,
			assets[i].FileSize, assets[i].FileFormat, assets[i].FileHashSha256,
		)
	}
	// 写入库
	insertSQL := query + strings.Join(valueStatement, ",")
	if err = m.Exec(insertSQL, values...).Error; err != nil {
		return err
	}
	return nil
}
