package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sjk1101/mySQL"
	"github.com/sjk1101/mySQL/example/mysql/model"
)

type modelDao struct {
	mySQL.MysqlBaseDAO
	model.BaseContext
	data  *model.Model
	query *ModelQuery
}

type ModelQuery struct {
	ID        int    `uri:"id"`
	FieldLike string `json:"fieldLike"`
}

func NewModelDAO(b *model.BaseContext) *modelDao {
	return &modelDao{
		BaseContext: *b,
		data:        &model.Model{},
	}
}

func (d *modelDao) Where(db *gorm.DB) (*gorm.DB, error) {
	query := d.query
	if query != nil {
		if query.ID != 0 {
			db.Where("id = ?", query.ID)
		}
		if query.FieldLike != "" {
			db = db.Where("field1 LIKE ?", query.FieldLike)
		}
	}
	return db, nil
}

func (d *modelDao) GetModel() mySQL.MysqlModel {
	return d.data
}

func (d *modelDao) SetQuery(query *ModelQuery) *modelDao {
	d.query = query
	return d
}

func (d *modelDao) SetData(data *model.Model) *modelDao {
	d.data = data
	return d
}

func (d *modelDao) Get() (result *model.Model, err error) {
	result = &model.Model{}
	err = _mySQL.Get(d, result)
	if result.ID == 0 {
		return nil, err
	}
	return
}

func (d *modelDao) List() (result []*model.Model, pager *mySQL.Page, err error) {
	result = []*model.Model{}
	pager, err = _mySQL.List(d, &result)
	return
}

func (d *modelDao) Insert() (insertedID int, err error) {
	return _mySQL.Insert(d)
}

func (d *modelDao) Update() (rowsAffected int64, err error) {
	return _mySQL.Update(d)
}

func (d *modelDao) Delete() (rowsAffected int64, err error) {
	return _mySQL.Delete(d)
}
