// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTenantOrder is the golang structure of table hg_addon_hgexample_tenant_order for DAO operations like Where/Data.
type AddonHgexampleTenantOrder struct {
	g.Meta      `orm:"table:hg_addon_hgexample_tenant_order, do:true"`
	Id          interface{} // 主键
	TenantId    interface{} // 租户ID
	MerchantId  interface{} // 商户ID
	UserId      interface{} // 用户ID
	ProductName interface{} // 购买产品
	OrderSn     interface{} // 订单号
	Money       interface{} // 充值金额
	Remark      interface{} // 备注
	Status      interface{} // 订单状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
