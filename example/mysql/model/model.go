package model

import (
	"mySQL"
)

type Model struct {
	mySQL.MysqlBase
	Field1 string `gorm:"column:field1" json:"field1"`
}

// TableName sets the insert table Name for this struct type
func (m *Model) TableName() string {
	return "model"
}
