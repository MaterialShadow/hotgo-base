// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.13.1
package sysin

import (
	"context"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CurdDemoUpdateFields 修改CURD列表字段过滤
type CurdDemoUpdateFields struct {
	Title       string `json:"title"       dc:"标题"`
	Description string `json:"description" dc:"描述"`
	Content     string `json:"content"     dc:"内容"`
	Image       string `json:"image"       dc:"单图"`
	Attachfile  string `json:"attachfile"  dc:"附件"`
	CityId      int64  `json:"cityId"      dc:"所在城市"`
	Sort        int    `json:"sort"        dc:"排序"`
	Switch      int    `json:"switch"      dc:"显示开关"`
	Status      int    `json:"status"      dc:"状态"`
	UpdatedBy   int64  `json:"updatedBy"   dc:"更新者"`
	CategoryId  int64  `json:"categoryId"  dc:"测试分类"`
}

// CurdDemoInsertFields 新增CURD列表字段过滤
type CurdDemoInsertFields struct {
	Title       string `json:"title"       dc:"标题"`
	Description string `json:"description" dc:"描述"`
	Content     string `json:"content"     dc:"内容"`
	Image       string `json:"image"       dc:"单图"`
	Attachfile  string `json:"attachfile"  dc:"附件"`
	CityId      int64  `json:"cityId"      dc:"所在城市"`
	Sort        int    `json:"sort"        dc:"排序"`
	Switch      int    `json:"switch"      dc:"显示开关"`
	Status      int    `json:"status"      dc:"状态"`
	CreatedBy   int64  `json:"createdBy"   dc:"创建者"`
	CategoryId  int64  `json:"categoryId"  dc:"测试分类"`
}

// CurdDemoEditInp 修改/新增CURD列表
type CurdDemoEditInp struct {
	entity.SysGenCurdDemo
}

func (in *CurdDemoEditInp) Filter(ctx context.Context) (err error) {
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证描述
	if err := g.Validator().Rules("required").Data(in.Description).Messages("描述不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证测试分类
	if err := g.Validator().Rules("required").Data(in.CategoryId).Messages("测试分类不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type CurdDemoEditModel struct{}

// CurdDemoDeleteInp 删除CURD列表
type CurdDemoDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *CurdDemoDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoDeleteModel struct{}

// CurdDemoViewInp 获取指定CURD列表信息
type CurdDemoViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *CurdDemoViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoViewModel struct {
	entity.SysGenCurdDemo
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"更新者摘要信息"`
}

// CurdDemoListInp 获取CURD列表列表
type CurdDemoListInp struct {
	form.PageReq
	Id               int64         `json:"id"               dc:"ID"`
	Title            string        `json:"title"            dc:"标题"`
	Description      string        `json:"description"      dc:"描述"`
	Status           int           `json:"status"           dc:"状态"`
	CreatedBy        string        `json:"createdBy"        dc:"创建者"`
	CreatedAt        []*gtime.Time `json:"createdAt"        dc:"创建时间"`
	CategoryId       int64         `json:"categoryId"       dc:"测试分类"`
	TestCategoryName string        `json:"testCategoryName" dc:"关联分类"`
}

func (in *CurdDemoListInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoListModel struct {
	Id               int64             `json:"id"               dc:"ID"`
	Title            string            `json:"title"            dc:"标题"`
	Description      string            `json:"description"      dc:"描述"`
	Image            string            `json:"image"            dc:"单图"`
	Attachfile       string            `json:"attachfile"       dc:"附件"`
	Sort             int               `json:"sort"             dc:"排序"`
	Switch           int               `json:"switch"           dc:"显示开关"`
	Status           int               `json:"status"           dc:"状态"`
	CreatedBy        int64             `json:"createdBy"        dc:"创建者"`
	CreatedBySumma   *hook.MemberSumma `json:"createdBySumma"   dc:"创建者摘要信息"`
	CreatedAt        *gtime.Time       `json:"createdAt"        dc:"创建时间"`
	UpdatedBy        int64             `json:"updatedBy"        dc:"更新者"`
	UpdatedBySumma   *hook.MemberSumma `json:"updatedBySumma"   dc:"更新者摘要信息"`
	UpdatedAt        *gtime.Time       `json:"updatedAt"        dc:"修改时间"`
	CategoryId       int64             `json:"categoryId"       dc:"测试分类"`
	TestCategoryName string            `json:"testCategoryName" dc:"关联分类"`
}

// CurdDemoExportModel 导出CURD列表
type CurdDemoExportModel struct {
	Id               int64       `json:"id"               dc:"ID"`
	Title            string      `json:"title"            dc:"标题"`
	Description      string      `json:"description"      dc:"描述"`
	Image            string      `json:"image"            dc:"单图"`
	Attachfile       string      `json:"attachfile"       dc:"附件"`
	CityId           int64       `json:"cityId"           dc:"所在城市"`
	Sort             int         `json:"sort"             dc:"排序"`
	Switch           int         `json:"switch"           dc:"显示开关"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedBy        int64       `json:"createdBy"        dc:"创建者"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedBy        int64       `json:"updatedBy"        dc:"更新者"`
	CategoryId       int64       `json:"categoryId"       dc:"测试分类"`
	TestCategoryName string      `json:"testCategoryName" dc:"关联分类"`
}

// CurdDemoMaxSortInp 获取CURD列表最大排序
type CurdDemoMaxSortInp struct{}

func (in *CurdDemoMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CurdDemoSwitchInp 更新CURD列表开关状态
type CurdDemoSwitchInp struct {
	form.SwitchReq
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *CurdDemoSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoSwitchModel struct{}
