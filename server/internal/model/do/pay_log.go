// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayLog is the golang structure of table hg_pay_log for DAO operations like Where/Data.
type PayLog struct {
	g.Meta        `orm:"table:hg_pay_log, do:true"`
	Id            interface{} // 主键
	MemberId      interface{} // 会员ID
	AppId         interface{} // 应用ID
	AddonsName    interface{} // 插件名称
	OrderSn       interface{} // 关联订单号
	OrderGroup    interface{} // 组别[默认统一支付类型]
	Openid        interface{} // openid
	MchId         interface{} // 商户支付账户
	Subject       interface{} // 订单标题
	Detail        *gjson.Json // 支付商品详情
	AuthCode      interface{} // 刷卡码
	OutTradeNo    interface{} // 商户订单号
	TransactionId interface{} // 交易号
	PayType       interface{} // 支付类型
	PayAmount     interface{} // 支付金额
	ActualAmount  interface{} // 实付金额
	PayStatus     interface{} // 支付状态
	PayAt         *gtime.Time // 支付时间
	TradeType     interface{} // 交易类型
	RefundSn      interface{} // 退款单号
	IsRefund      interface{} // 是否退款
	Custom        interface{} // 自定义参数
	CreateIp      interface{} // 创建者IP
	PayIp         interface{} // 支付者IP
	NotifyUrl     interface{} // 支付通知回调地址
	ReturnUrl     interface{} // 买家付款成功跳转地址
	TraceIds      *gjson.Json // 链路ID集合
	Status        interface{} // 状态
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 修改时间
}
