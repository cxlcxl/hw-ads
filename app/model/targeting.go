package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Targeting struct {
	connectDb

	Id                 int64     `db:"id"`
	AccountId          int64     `db:"account_id"`
	AdvertiserId       string    `db:"advertiser_id"` // 广告主账户ID
	TargetingId        int64     `db:"targeting_id"`
	TargetingName      string    `db:"targeting_name"`        // 定向名称
	TargetingType      string    `db:"targeting_type"`        // 定向类型
	LocationType       string    `db:"location_type"`         // 地域定向类型
	IncludeLocation    string    `db:"include_location"`      // 地域 - 包含
	ExcludeLocation    string    `db:"exclude_location"`      // 地域 - 排除
	Carriers           string    `db:"carriers"`              // 运营商
	Language           string    `db:"language"`              // 语言
	Age                string    `db:"age"`                   // 年龄
	Gender             string    `db:"gender"`                // 性别
	AppCategory        string    `db:"app_category"`          // App 行为类型
	AppCategories      string    `db:"app_categories"`        // App 行为
	InstalledApps      string    `db:"installed_apps"`        // app 安装
	AppInterest        string    `db:"app_interest"`          // App 兴趣类型
	AppInterests       string    `db:"app_interests"`         // App 兴趣
	Series             string    `db:"series"`                // 设备
	NetworkType        string    `db:"network_type"`          // 联网方式
	NotAudiences       string    `db:"not_audiences"`         // 排除人群
	Audiences          string    `db:"audiences"`             // 包含人群
	AppCategoryOfMedia string    `db:"app_category_of_media"` // 媒体类型
	CreatedAt          time.Time `db:"created_at"`            // 添加时间
	UpdatedAt          time.Time `db:"updated_at"`            // 最后一次修改时间
}

func NewTargeting(db *gorm.DB) *Targeting {
	return &Targeting{connectDb: connectDb{DB: db}}
}

func (m *Targeting) TableName() string {
	return "targetings"
}

func (m *Targeting) GetTargets(accountId int64) (targets []*Targeting, err error) {
	query := m.Table(m.TableName()).Order("id desc")
	if accountId > 0 {
		query = query.Where("account_id = ?", accountId)
	}
	err = query.Find(&targets).Error
	return
}

// CheckExistsByTargetName 名称检查定向包是否存在
func (m *Targeting) CheckExistsByTargetName(targetName string) (total int64, err error) {
	err = m.Table(m.TableName()).Where("targeting_name = ?", targetName).Limit(1).Count(&total).Error
	return
}

func (m *Targeting) TargetingCreate(targeting *Targeting) (err error) {
	err = m.Table(m.TableName()).Create(targeting).Error
	return
}

func (m *Targeting) BatchInsert(targetings []*Targeting) (err error) {
	rows := "account_id,advertiser_id,targeting_id,targeting_name,targeting_type,location_type,include_location," +
		"exclude_location,carriers,language,age,gender,app_category,app_categories,installed_apps,app_interest," +
		"app_interests,series,network_type,not_audiences,audiences,app_category_of_media,created_at,updated_at"
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.TableName(), rows)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	targetingIds := make([]int64, 0)
	for _, targeting := range targetings {
		targetingIds = append(targetingIds, targeting.TargetingId)
	}
	err = m.Transaction(func(session *gorm.DB) error {
		deleteQuery := fmt.Sprintf("delete from %s where targeting_id in ?", m.TableName())
		if err = session.Exec(deleteQuery, targetingIds).Error; err != nil {
			return err
		}

		for i := 0; i < len(targetings); i++ {
			data := targetings[i]
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				data.AccountId, data.AdvertiserId, data.TargetingId, data.TargetingName, data.TargetingType,
				data.LocationType, data.IncludeLocation, data.ExcludeLocation, data.Carriers, data.Language,
				data.Age, data.Gender, data.AppCategory, data.AppCategories, data.InstalledApps, data.AppInterest,
				data.AppInterests, data.Series, data.NetworkType, data.NotAudiences, data.Audiences,
				data.AppCategoryOfMedia, data.CreatedAt, data.UpdatedAt,
			)
			// 达到了 500 条数据，或最后一条了
			if chunk == 500 || i == len(targetings)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",")
				if err = session.Exec(insertSQL, values...).Error; err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}
		return nil
	})
	return err
}
