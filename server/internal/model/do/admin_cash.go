// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCash is the golang structure of table hg_admin_cash for DAO operations like Where/Data.
type AdminCash struct {
	g.Meta    `orm:"table:hg_admin_cash, do:true"`
	Id        interface{} // ID
	MemberId  interface{} // 管理员ID
	Money     interface{} // 提现金额
	Fee       interface{} // 手续费
	LastMoney interface{} // 最终到账金额
	Ip        interface{} // 申请人IP
	Status    interface{} // 状态码
	Msg       interface{} // 处理结果
	HandleAt  *gtime.Time // 处理时间
	CreatedAt *gtime.Time // 申请时间
}
