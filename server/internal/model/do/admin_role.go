// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure of table hg_admin_role for DAO operations like Where/Data.
type AdminRole struct {
	g.Meta     `orm:"table:hg_admin_role, do:true"`
	Id         interface{} // 角色ID
	Name       interface{} // 角色名称
	Key        interface{} // 角色权限字符串
	DataScope  interface{} // 数据范围
	CustomDept *gjson.Json // 自定义部门权限
	Pid        interface{} // 上级角色ID
	Level      interface{} // 关系树等级
	Tree       interface{} // 关系树
	Remark     interface{} // 备注
	Sort       interface{} // 排序
	Status     interface{} // 角色状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
