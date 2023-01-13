package model

import (
	"bs.mobgi.cc/app/cache"
	"gorm.io/gorm"
)

type Continent struct {
	connectDb
	Id    int64  `json:"id"`
	CName string `json:"c_name"`
}

func (m *Continent) TableName() string {
	return "continents"
}

var (
	continentsKey = "db:continents"
)

func NewContinent(db *gorm.DB) *Continent {
	return &Continent{connectDb: connectDb{DB: db}}
}
func (m *Continent) Continents() (continents []*Continent, err error) {
	err = cache.New(m.DB).Query(continentsKey, &continents, func(db *gorm.DB, v interface{}) error {
		return db.Table(m.TableName()).Find(v).Error
	})
	return
}
