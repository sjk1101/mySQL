package mySQL

import (
	"math"
)

type Page struct {
	Index int64 `json:"index" form:"pi"`
	Size  int64 `json:"size" form:"ps"`
	Total int64 `json:"total"`
	Count int64 `json:"pages"`
}

func (p *Page) MakePageInfo(count int64) {
	p.Total = count
	p.Count = int64(math.Ceil(float64(count) / float64(p.Size)))
}
