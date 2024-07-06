// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDept is the golang structure of table hg_admin_dept for DAO operations like Where/Data.
type AdminDept struct {
	g.Meta    `orm:"table:hg_admin_dept, do:true"`
	Id        interface{} // 部门ID
	Pid       interface{} // 父部门ID
	Name      interface{} // 部门名称
	Code      interface{} // 部门编码
	Type      interface{} // 部门类型
	Leader    interface{} // 负责人
	Phone     interface{} // 联系电话
	Email     interface{} // 邮箱
	Level     interface{} // 关系树等级
	Tree      interface{} // 关系树
	Sort      interface{} // 排序
	Status    interface{} // 部门状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
