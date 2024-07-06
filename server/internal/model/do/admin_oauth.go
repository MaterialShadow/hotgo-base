// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOauth is the golang structure of table hg_admin_oauth for DAO operations like Where/Data.
type AdminOauth struct {
	g.Meta       `orm:"table:hg_admin_oauth, do:true"`
	Id           interface{} // 主键
	MemberId     interface{} // 用户ID
	Unionid      interface{} // 唯一ID
	OauthClient  interface{} // 授权组别
	OauthOpenid  interface{} // 授权开放ID
	Sex          interface{} // 性别
	Nickname     interface{} // 昵称
	HeadPortrait interface{} // 头像
	Birthday     *gtime.Time // 生日
	Country      interface{} // 国家
	Province     interface{} // 省
	City         interface{} // 市
	Status       interface{} // 状态
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 修改时间
}
