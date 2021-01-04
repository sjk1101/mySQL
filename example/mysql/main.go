package main

import (
	"mySQL/example/mysql/dao"
	"mySQL/example/mysql/model"
)

func main() {
	dao.InitMySQL()
	b := &model.BaseContext{}
	b.PageSize = 10
	b.PageIndex = 1
	_, _, _ = dao.NewModelDAO(b).SetQuery(&dao.ModelQuery{
		ID: 1,
	}).List()
}
