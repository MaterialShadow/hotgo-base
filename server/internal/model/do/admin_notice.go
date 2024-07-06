// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNotice is the golang structure of table hg_admin_notice for DAO operations like Where/Data.
type AdminNotice struct {
	g.Meta    `orm:"table:hg_admin_notice, do:true"`
	Id        interface{} // 公告ID
	Title     interface{} // 公告标题
	Type      interface{} // 公告类型
	Tag       interface{} // 标签
	Content   interface{} // 公告内容
	Receiver  *gjson.Json // 接收者
	Remark    interface{} // 备注
	Sort      interface{} // 排序
	Status    interface{} // 公告状态
	CreatedBy interface{} // 发送人
	UpdatedBy interface{} // 修改人
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
