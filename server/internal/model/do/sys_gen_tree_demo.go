// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTreeDemo is the golang structure of table hg_sys_gen_tree_demo for DAO operations like Where/Data.
type SysGenTreeDemo struct {
	g.Meta      `orm:"table:hg_sys_gen_tree_demo, do:true"`
	Id          interface{} // ID
	Pid         interface{} // 上级ID
	Level       interface{} // 关系树级别
	Tree        interface{} // 关系树
	CategoryId  interface{} // 分类ID
	Title       interface{} // 标题
	Description interface{} // 描述
	Sort        interface{} // 排序
	Status      interface{} // 状态
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
