// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLog is the golang structure of table hg_sys_sms_log for DAO operations like Where/Data.
type SysSmsLog struct {
	g.Meta    `orm:"table:hg_sys_sms_log, do:true"`
	Id        interface{} // 主键
	Event     interface{} // 事件
	Mobile    interface{} // 手机号
	Code      interface{} // 验证码或短信内容
	Times     interface{} // 验证次数
	Ip        interface{} // ip地址
	Status    interface{} // 状态(1未验证,2已验证)
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
