package mySQL

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	FuzzyMode  = "fuzzy"
	PrefixMode = "prefix"
)

type QueryMode string

func (q *QueryMode) SetQuery(db *gorm.DB, column string, value interface{}) *gorm.DB {
	if *q == FuzzyMode {
		db = db.Where(column+" LIKE ?", fmt.Sprintf("%%%v%%", value))
	} else if *q == PrefixMode {
		db = db.Where(column+" LIKE ?", fmt.Sprintf("%v%%", value))
	} else {
		db = db.Where(column+" = ?", value)
	}
	return db

}
