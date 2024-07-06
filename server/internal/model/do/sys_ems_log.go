// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysEmsLog is the golang structure of table hg_sys_ems_log for DAO operations like Where/Data.
type SysEmsLog struct {
	g.Meta    `orm:"table:hg_sys_ems_log, do:true"`
	Id        interface{} // 主键
	Event     interface{} // 事件
	Email     interface{} // 邮箱地址，多个用;隔开
	Code      interface{} // 验证码
	Times     interface{} // 验证次数
	Content   interface{} // 邮件内容
	Ip        interface{} // ip地址
	Status    interface{} // 状态(1未验证,2已验证)
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
