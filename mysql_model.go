package mySQL

import (
	"time"
)

type MysqlModel interface {
	TableName() string
}

type IDGetter interface {
	GetID() int
}

type DefaultSortsGetter interface {
	GetDefaultSorts() Sorts
}

type UserInfoSetter interface {
	SetCreator(user string)
	SetEditor(user string)
}

type MysqlBase struct {
	MysqlBaseWithoutSort
	Sort int `gorm:"column:sort" json:"sort"`
}

type MysqlBaseWithoutSort struct {
	MysqlBaseOnlyID
	Enable      *int        `gorm:"column:enable" json:"enable"`
	Description string `gorm:"column:description" json:"description"`
}

type MysqlBaseOnlyID struct {
	ID        int       `gorm:"column:id;primary_key;auto_increment" json:"id" form:"id" uri:"id"`
	Creator   string    `gorm:"column:creator" json:"creator"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	Editor    string    `gorm:"column:editor" json:"editor"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

var _ IDGetter = (*MysqlBase)(nil)
var _ DefaultSortsGetter = (*MysqlBase)(nil)
var _ UserInfoSetter = (*MysqlBase)(nil)

func (m *MysqlBaseOnlyID) GetID() int {
	return m.ID
}

func (m *MysqlBaseOnlyID) SetCreator(user string) {
	m.Creator = user
}

func (m *MysqlBaseOnlyID) SetEditor(user string) {
	m.Editor = user
}

func (m *MysqlBase) GetDefaultSorts() Sorts {
	return Sorts{{
		Column: "sort",
		IsDesc: false,
	}}
}
