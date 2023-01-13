package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Tracking struct {
	connectDb

	Id               int64  `json:"id"`
	AccountId        int64  `json:"account_id"`         // 账户ID;对应 accounts 表的id字段
	AdvertiserId     string `json:"advertiser_id"`      // 广告主账户ID
	AppId            string `json:"app_id"`             // 第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位
	TrackingId       int64  `json:"tracking_id"`        // 转化跟踪指标ID
	EffectType       string `json:"effect_type"`        // 转化目标
	EffectName       string `json:"effect_name"`        // 转化跟踪指标名称
	ClickTrackingUrl string `json:"click_tracking_url"` // 点击监测地址
	ImpTrackingUrl   string `json:"imp_tracking_url"`   // 曝光监测地址
}

func (m *Tracking) TableName() string {
	return "trackings"
}

func NewTracking(db *gorm.DB) *Tracking {
	return &Tracking{connectDb: connectDb{DB: db}}
}

func (m *Tracking) TrackingList(appId string) (trackings []*Tracking, err error) {
	err = m.Table(m.TableName()).
		Select("id", "account_id", "advertiser_id", "app_id", "tracking_id", "effect_type", "effect_name", "'' as click_tracking_url", "'' as imp_tracking_url").
		Where("app_id = ?", appId).Find(&trackings).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (m *Tracking) BatchInsert(tracks []*Tracking, disabled []int64) error {
	trackingRows := "account_id,advertiser_id,app_id,tracking_id,effect_type,effect_name,click_tracking_url,imp_tracking_url"
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.TableName(), trackingRows)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	for i := 0; i < len(tracks); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?)")
		values = append(values,
			tracks[i].AccountId, tracks[i].AdvertiserId, tracks[i].AppId, tracks[i].TrackingId,
			tracks[i].EffectType, tracks[i].EffectName, tracks[i].ClickTrackingUrl, tracks[i].ImpTrackingUrl,
		)
	}
	// 写入库
	insertSQL := query + strings.Join(valueStatement, ",")
	return m.Transaction(func(tx *gorm.DB) error {
		if len(disabled) > 0 {
			if err := tx.Exec(
				fmt.Sprintf("delete from %s where app_id = ? and tracking_id in ?", m.TableName()),
				tracks[0].AppId,
				disabled,
			).Error; err != nil {
				return err
			}
		}
		if err := tx.Exec(insertSQL, values...).Error; err != nil {
			return err
		}
		return nil
	})
}
