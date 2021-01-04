package mySQL

import (
	"github.com/jinzhu/gorm"
)

type Preload struct {
	column     string
	conditions []interface{}
}

type MysqlDAO interface {
	Contextor
	Sorter
	GetDB() *gorm.DB
	GetModel() MysqlModel
	Where(db *gorm.DB) (*gorm.DB, error)
	GetPreloads() []*Preload
	SetCommandType(t CommandType)
}

type MysqlBaseDAO struct {
	db       *gorm.DB
	preloads []*Preload
	CommandTypeBase
	Sorts
}

func (d *MysqlBaseDAO) SetDB(db *gorm.DB) {
	d.db = db
}

func (d *MysqlBaseDAO) GetDB() *gorm.DB {
	return d.db
}

func (d *MysqlBaseDAO) Preload(column string, conditions ...interface{}) {
	d.preloads = append(d.preloads, &Preload{
		column:     column,
		conditions: conditions,
	})
}

func (d *MysqlBaseDAO) GetPreloads() []*Preload {
	return d.preloads
}
