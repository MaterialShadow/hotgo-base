// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCron is the golang structure of table hg_sys_cron for DAO operations like Where/Data.
type SysCron struct {
	g.Meta    `orm:"table:hg_sys_cron, do:true"`
	Id        interface{} // 任务ID
	GroupId   interface{} // 分组ID
	Title     interface{} // 任务标题
	Name      interface{} // 任务方法
	Params    interface{} // 函数参数
	Pattern   interface{} // 表达式
	Policy    interface{} // 策略
	Count     interface{} // 执行次数
	Sort      interface{} // 排序
	Remark    interface{} // 备注
	Status    interface{} // 任务状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
