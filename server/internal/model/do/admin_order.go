// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOrder is the golang structure of table hg_admin_order for DAO operations like Where/Data.
type AdminOrder struct {
	g.Meta             `orm:"table:hg_admin_order, do:true"`
	Id                 interface{} // 主键
	MemberId           interface{} // 管理员id
	OrderType          interface{} // 订单类型
	ProductId          interface{} // 产品id
	OrderSn            interface{} // 关联订单号
	Money              interface{} // 充值金额
	Remark             interface{} // 备注
	RefundReason       interface{} // 退款原因
	RejectRefundReason interface{} // 拒绝退款原因
	Status             interface{} // 状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
