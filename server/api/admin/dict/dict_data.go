// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// DataEditReq 修改/新增字典数据
type DataEditReq struct {
	g.Meta `path:"/dictData/edit" method:"post" tags:"字典数据" summary:"修改/新增字典数据"`
	sysin.DictDataEditInp
}

type DataEditRes struct{}

// DataDeleteReq 删除字典数据
type DataDeleteReq struct {
	g.Meta `path:"/dictData/delete" method:"post" tags:"字典数据" summary:"删除字典数据"`
	sysin.DictDataDeleteInp
}

type DataDeleteRes struct{}

// DataListReq 查询列表
type DataListReq struct {
	g.Meta `path:"/dictData/list" method:"get" tags:"字典数据" summary:"获取字典数据列表"`
	sysin.DictDataListInp
}

type DataListRes struct {
	List []*sysin.DictDataListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

type DataSelectReq struct {
	g.Meta `path:"/dictData/option/{Type}" method:"get" tags:"字典数据" summary:"获取指定字典选项"`
	sysin.DataSelectInp
}

type DataSelectRes sysin.DataSelectModel

type DataSelectsReq struct {
	g.Meta `path:"/dictData/options" method:"get" tags:"字典数据" summary:"获取多个字典选项"`
	Types  []string `json:"types"`
}

type DataSelectsRes map[string]sysin.DataSelectModel
