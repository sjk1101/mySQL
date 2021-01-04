package mySQL

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type MySQLConfig struct {
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

type MySQL struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
}

func NewMySQL(cfg MySQLConfig) (m *MySQL, err error) {
	m = &MySQL{}
	m.readDB = cfg.ReadDB
	m.writeDB = cfg.WriteDB

	if m.readDB == nil {
		err = errors.New("mysql read db connect is nil")
		return nil, err
	}

	if m.writeDB == nil {
		err = errors.New("mysql write db connect is nil")
		return nil, err
	}
	return m, nil
}

func (m *MySQL) applySorts(dao MysqlDAO, db *gorm.DB) *gorm.DB {
	sorts := dao.GetSorts()
	if len(sorts) == 0 {
		sortGetter, ok := dao.GetModel().(DefaultSortsGetter)
		if ok {
			sorts = sortGetter.GetDefaultSorts()
		}
	}

	for _, sort := range sorts {
		value := sort.Column
		if sort.IsDesc {
			value += " DESC"
		}
		db = db.Order(value)
	}
	return db
}

func (m *MySQL) ApplyPreloads(dao MysqlDAO, db *gorm.DB) *gorm.DB {
	for _, preload := range dao.GetPreloads() {
		db = db.Preload(
			preload.column,
			preload.conditions...,
		)
	}
	return db
}

func (m *MySQL) SetPager(ctx Contextor, db *gorm.DB) (*gorm.DB, *Page) {
	c := ctx.GetContext()
	if c.PageSize == 0 {
		return db, nil
	}
	page := &Page{}
	page.Index = c.PageIndex
	page.Size = c.PageSize
	limit := c.PageSize
	skip := (c.PageIndex - 1) * c.PageSize
	return db.Limit(limit).Offset(skip), page
}

func (m *MySQL) Get(dao MysqlDAO, item interface{}) (err error) {
	query := dao.GetDB()
	if query == nil {
		query = m.readDB
	}
	dao.SetCommandType(Get)

	query, err = dao.Where(query)
	if err != nil {
		return
	}

	query = m.ApplyPreloads(dao, query)
	query = m.applySorts(dao, query)
	err = query.Take(item).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return
}

func (m *MySQL) List(dao MysqlDAO, item interface{}) (pager *Page, err error) {
	query := dao.GetDB()
	if query == nil {
		query = m.readDB
	}
	dao.SetCommandType(List)

	query, err = dao.Where(query)
	if err != nil {
		return
	}

	query = m.ApplyPreloads(dao, query)
	query = m.applySorts(dao, query)
	query, pager = m.SetPager(dao, query)
	query = query.Find(item)
	err = query.Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if pager != nil {
		var count int64
		query.Offset(-1).Limit(-1).Count(&count)
		pager.MakePageInfo(count)
	}

	return
}

func (m *MySQL) Insert(dao MysqlDAO) (insertedID int, err error) {
	cmd := dao.GetDB()
	if cmd == nil {
		cmd = m.writeDB
	}
	dao.SetCommandType(Insert)
	data := dao.GetModel()
	if userContextor, ok := dao.(UserContextor); ok {
		if userInfoSetter, ok := data.(UserInfoSetter); ok {
			userInfoSetter.SetCreator(userContextor.GetUserName())
			userInfoSetter.SetEditor(userContextor.GetUserName())
		}
	}
	cmd = cmd.Create(data)
	err = cmd.Error
	if err != nil {
		return
	}

	val := cmd.Value
	if result, ok := val.(IDGetter); ok {
		insertedID = result.GetID()
	}
	return
}

func (m *MySQL) InsertIgnore(dao MysqlDAO) (insertedID int, rowsAffected int64, err error) {
	cmd := dao.GetDB()
	if cmd == nil {
		cmd = m.writeDB
	}
	dao.SetCommandType(Insert)
	data := dao.GetModel()
	if userContextor, ok := dao.(UserContextor); ok {
		if userInfoSetter, ok := data.(UserInfoSetter); ok {
			userInfoSetter.SetCreator(userContextor.GetUserName())
			userInfoSetter.SetEditor(userContextor.GetUserName())
		}
	}
	cmd = cmd.Set("gorm:insert_modifier", "IGNORE").Create(data)
	err = cmd.Error
	if err != nil {
		return
	}

	val := cmd.Value
	if result, ok := val.(IDGetter); ok {
		insertedID = result.GetID()
	}

	rowsAffected = cmd.RowsAffected
	return
}

func (m *MySQL) FirstOrCreate(dao MysqlDAO, item interface{}) (rowsAffected int64, err error) {
	db := dao.GetDB()
	if db == nil {
		db = m.writeDB
	}
	dao.SetCommandType(Insert)
	query := db
	query, err = dao.Where(query)
	if err != nil {
		return
	}

	err = query.Take(item).Error
	if err == nil {
		return
	}

	if !gorm.IsRecordNotFoundError(err) {
		return
	}
	err = nil

	cmd := db
	data := dao.GetModel()
	if userContextor, ok := dao.(UserContextor); ok {
		if userInfoSetter, ok := data.(UserInfoSetter); ok {
			userInfoSetter.SetCreator(userContextor.GetUserName())
			userInfoSetter.SetEditor(userContextor.GetUserName())
		}
	}
	cmd = cmd.Set("gorm:insert_modifier", "IGNORE").Create(data)
	err = cmd.Error
	if err != nil {
		return
	}
	rowsAffected = cmd.RowsAffected

	err = query.Take(item).Error
	return
}

func (m *MySQL) Update(dao MysqlDAO) (rowsAffected int64, err error) {
	cmd := dao.GetDB()
	if cmd == nil {
		cmd = m.writeDB
	}
	dao.SetCommandType(Update)
	cmd, err = dao.Where(cmd)
	if err != nil {
		return
	}
	data := dao.GetModel()
	if userContextor, ok := dao.(UserContextor); ok {
		if userInfoSetter, ok := data.(UserInfoSetter); ok {
			userInfoSetter.SetEditor(userContextor.GetUserName())
		}
	}
	cmd = cmd.Model(dao.GetModel()).Updates(data)
	err = cmd.Error
	if err != nil {
		return
	}

	rowsAffected = cmd.RowsAffected
	return
}

func (m *MySQL) UpdateSort(dao MysqlDAO) (rowsAffected int64, err error) {
	cmd := dao.GetDB()
	if cmd == nil {
		cmd = m.writeDB
	}
	dao.SetCommandType(Update)
	data := dao.GetModel()
	if userContextor, ok := dao.(UserContextor); ok {
		if userInfoSetter, ok := data.(UserInfoSetter); ok {
			userInfoSetter.SetEditor(userContextor.GetUserName())
		}
	}
	cmd = cmd.Model(dao.GetModel()).Select("sort").Updates(data)
	err = cmd.Error
	if err != nil {
		return
	}

	rowsAffected = cmd.RowsAffected
	return
}

func (m *MySQL) Delete(dao MysqlDAO) (rowsAffected int64, err error) {
	cmd := dao.GetDB()
	if cmd == nil {
		cmd = m.writeDB
	}
	dao.SetCommandType(Delete)
	cmd, err = dao.Where(cmd)
	if err != nil {
		return
	}
	cmd = cmd.Delete(dao.GetModel())
	err = cmd.Error
	if err != nil {
		return
	}

	rowsAffected = cmd.RowsAffected
	return
}

func (m *MySQL) DeleteMany(dao MysqlDAO) (rowsAffected int64, err error) {
	cmd := dao.GetDB()
	if cmd == nil {
		cmd = m.writeDB
	}
	dao.SetCommandType(Delete)
	cmd, err = dao.Where(cmd)
	if err != nil {
		return
	}
	cmd = cmd.Delete(dao.GetModel())
	err = cmd.Error
	if err != nil {
		return
	}

	rowsAffected = cmd.RowsAffected
	return
}
