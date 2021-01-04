package mySQL

import (
	"github.com/gin-gonic/gin"
	"mySQL/core"
)

type Contextor interface {
	GetContext() Context
}

type UserContextor interface {
	GetUserName() string
}

type Context struct {
	PageIndex int64 `form:"pi"`
	PageSize  int64 `form:"ps"`
}

func (c Context) GetContext() Context {
	return c
}

func (c Context) ToCoreContext() *core.Context {
	ctx := &core.Context{}
	if c.PageSize > 0 {
		ctx.SetItem("ps", c.PageSize)
		ctx.SetItem("pi", c.PageIndex)
	}
	return ctx
}

func InitContext(c *gin.Context) Context {
	baseContext := Context{}
	err := c.ShouldBindQuery(&baseContext)
	if err != nil {
		return baseContext
	}

	return baseContext
}
