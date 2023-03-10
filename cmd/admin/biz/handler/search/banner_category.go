package search

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/searches"
	"github.com/quarkcms/quark-micro/cmd/admin/biz/model"
	"gorm.io/gorm"
)

type BannerCategory struct {
	searches.Select
}

// 初始化
func (p *BannerCategory) Init(column string, name string) *BannerCategory {
	p.ParentInit()

	p.Column = column
	p.Name = name

	return p
}

// 执行查询
func (p *BannerCategory) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	if value == "" || value == nil {
		return query
	}

	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *BannerCategory) Options(ctx *builder.Context) map[interface{}]interface{} {
	options, _ := (&model.BannerCategory{}).Options()

	return options
}
