package model

import (
	"mySQL"
)

type BaseContext struct {
	mySQL.Context
	TestField      string
	UserName       string
	UserPermission int
}
