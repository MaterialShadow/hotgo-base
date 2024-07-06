// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBlacklist is the golang structure of table hg_sys_blacklist for DAO operations like Where/Data.
type SysBlacklist struct {
	g.Meta    `orm:"table:hg_sys_blacklist, do:true"`
	Id        interface{} // 黑名单ID
	Ip        interface{} // IP地址
	Remark    interface{} // 备注
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
