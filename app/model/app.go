package model

import (
	"bs.mobgi.cc/app/cache"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type App struct {
	connectDb

	Id         int64   `json:"id"`
	AppId      string  `json:"app_id"`     // 第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位
	AppName    string  `json:"app_name"`   // 应用名称
	PkgName    string  `json:"pkg_name"`   // 应用包名或BundleID
	Channel    int64   `json:"channel"`    // 系统平台(渠道)：华为 AppGallery；GooglePlay; AppStore
	Tags       string  `json:"tags"`       // 应用标签
	IconUrl    string  `json:"icon_url"`   // 图标
	ProductId  string  `json:"product_id"` // 产品ID，创建任务时需要
	AccountIds []int64 `json:"account_ids" gorm:"-"`

	Timestamp
}

type SimpleApp struct {
	AppId   string `json:"app_id"`
	AppName string `json:"app_name"`
}

var (
	appsKey = "db:apps"
)

func (m *App) TableName() string {
	return "apps"
}

func NewApp(db *gorm.DB) *App {
	return &App{connectDb: connectDb{DB: db}}
}

func (m *App) AppList(appId, appName string, channel int64, actIds []int64, offset, limit int64) (apps []*App, total int64, err error) {
	query := m.Table(m.TableName()).Order("id desc")
	if len(appId) > 0 {
		query = query.Where("app_id like ?", "%"+appId+"%")
	}
	if len(appName) > 0 {
		query = query.Where("app_name like ?", "%"+appName+"%")
	}
	if channel > 0 {
		query = query.Where("channel = ?", channel)
	}
	if len(actIds) > 0 {
		in := fmt.Sprintf("app_id in (select app_id from `%s` where account_id in ?)", NewAppAct(nil).TableName())
		query = query.Where(in, actIds)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = query.Offset(int(offset)).Limit(int(limit)).Find(&apps).Error
	}
	return
}

func (m *App) AppCampaignList(appName string, offset, limit int64) (apps []*App, total int64, err error) {
	query := m.Table(m.TableName()).Order("id desc").
		Select("id", "app_id", "app_name", "icon_url")
	if len(appName) > 0 {
		query = query.Where("app_name like ? or app_id = ?", "%"+appName+"%", appName)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = query.Offset(int(offset)).Limit(int(limit)).Find(&apps).Error
	}
	return
}

func (m *App) FindAppByAppId(appId string) (app *App, err error) {
	err = m.Table(m.TableName()).Where("app_id = ?", appId).First(&app).Error
	return
}

func (m *App) FindAppById(id int64) (app *App, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&app).Error
	return
}

func (m *App) AllApps(actIds []int64) (apps []*SimpleApp, err error) {
	if len(actIds) > 0 {
		in := fmt.Sprintf("app_id in (select app_id from `%s` where account_id in ?)", NewAppAct(nil).TableName())
		err = m.Table(m.TableName()).
			Where(in, actIds).
			Select("app_id,app_name").Find(&apps).Error
	} else {
		err = cache.New(m.DB).Query(appsKey, &apps, func(db *gorm.DB, v interface{}) error {
			return db.Table(m.TableName()).Select("app_id,app_name").Find(v).Error
		})
	}
	return
}

func (m *App) CreateApp(app *App) (err error) {
	err = m.Table(m.TableName()).Create(app).Error
	return
}

func (m *App) UpdateApp(d map[string]interface{}, id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ? ", id).Updates(d).Error
	if err == nil {
		_ = cache.New(nil).DelQueryRowCache(appsKey, "")
	}
	return
}

type JobApp struct {
	SimpleApp
	PkgName string `json:"pkg_name"`
}

func (m *App) JobGetApps() (apps []*JobApp, err error) {
	err = m.Table(m.TableName()).Select("app_id", "app_name", "pkg_name").Find(&apps).Error
	return
}

func (m *App) BatchInsert(apps []*App, appActs []*AppAccount) (err error) {
	if len(apps) == 0 {
		return nil
	}
	updateColumns := []string{"app_name", "product_id", "pkg_name", "icon_url"}
	return m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns(updateColumns),
		}).CreateInBatches(apps, 100).Error; err != nil {
			return err
		}

		if err = NewAppAct(tx).BatchInsert(appActs); err != nil {
			return err
		}
		return nil
	})
}
