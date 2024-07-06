// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMember is the golang structure of table hg_admin_member for DAO operations like Where/Data.
type AdminMember struct {
	g.Meta             `orm:"table:hg_admin_member, do:true"`
	Id                 interface{} // 管理员ID
	DeptId             interface{} // 部门ID
	RoleId             interface{} // 角色ID
	RealName           interface{} // 真实姓名
	Username           interface{} // 帐号
	PasswordHash       interface{} // 密码
	Salt               interface{} // 密码盐
	PasswordResetToken interface{} // 密码重置令牌
	Integral           interface{} // 积分
	Balance            interface{} // 余额
	Avatar             interface{} // 头像
	Sex                interface{} // 性别
	Qq                 interface{} // qq
	Email              interface{} // 邮箱
	Mobile             interface{} // 手机号码
	Birthday           *gtime.Time // 生日
	CityId             interface{} // 城市编码
	Address            interface{} // 联系地址
	Pid                interface{} // 上级管理员ID
	Level              interface{} // 关系树等级
	Tree               interface{} // 关系树
	InviteCode         interface{} // 邀请码
	Cash               *gjson.Json // 提现配置
	LastActiveAt       *gtime.Time // 最后活跃时间
	Remark             interface{} // 备注
	Status             interface{} // 状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
