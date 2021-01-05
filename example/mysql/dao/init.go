package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sjk1101/mySQL"
)

var _mySQL *mySQL.MySQL

func MySQL() *mySQL.MySQL {
	return _mySQL
}

func InitMySQL() {
	var err error
	_mySQL, err = mySQL.NewMySQL(mySQL.MySQLConfig{
		ReadDB:  &gorm.DB{},
		WriteDB: &gorm.DB{},
	})
	if err != nil {
		panic(err)
	}
}
