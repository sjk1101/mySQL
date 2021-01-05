package model

import (
	"github.com/sjk1101/mySQL"
)

type BaseContext struct {
	mySQL.Context
	TestField      string
	UserName       string
	UserPermission int
}
