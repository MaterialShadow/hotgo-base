// Package member
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package member

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// UpdateCashReq 更新提现信息
type UpdateCashReq struct {
	g.Meta `path:"/member/updateCash" method:"post" tags:"用户" summary:"更新提现信息"`
	adminin.MemberUpdateCashInp
}

type UpdateCashRes struct{}

// UpdateEmailReq 换绑邮箱
type UpdateEmailReq struct {
	g.Meta `path:"/member/updateEmail" method:"post" tags:"用户" summary:"换绑邮箱"`
	adminin.MemberUpdateEmailInp
}

type UpdateEmailRes struct{}

// UpdateMobileReq 换绑手机号
type UpdateMobileReq struct {
	g.Meta `path:"/member/updateMobile" method:"post" tags:"用户" summary:"换绑手机号"`
	adminin.MemberUpdateMobileInp
}

type UpdateMobileRes struct{}

// UpdateProfileReq 更新用户资料
type UpdateProfileReq struct {
	g.Meta `path:"/member/updateProfile" method:"post" tags:"用户" summary:"更新用户资料"`
	adminin.MemberUpdateProfileInp
}

type UpdateProfileRes struct{}

// UpdatePwdReq 修改登录密码
type UpdatePwdReq struct {
	g.Meta `path:"/member/updatePwd" method:"post" tags:"用户" summary:"重置密码"`
	adminin.MemberUpdatePwdInp
}

type UpdatePwdRes struct{}

// ResetPwdReq 重置密码
type ResetPwdReq struct {
	g.Meta `path:"/member/resetPwd" method:"post" tags:"用户" summary:"重置密码"`
	adminin.MemberResetPwdInp
}

type ResetPwdRes struct{}

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/member/list" method:"get" tags:"用户" summary:"获取用户列表"`
	adminin.MemberListInp
}

type ListRes struct {
	List []*adminin.MemberListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/member/view" method:"get" tags:"用户" summary:"获取指定信息"`
	adminin.MemberViewInp
}

type ViewRes struct {
	*adminin.MemberViewModel
}

// EditReq 修改/新增
type EditReq struct {
	g.Meta `path:"/member/edit" method:"post" tags:"用户" summary:"修改/新增用户"`
	adminin.MemberEditInp
}

type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/member/delete" method:"post" tags:"用户" summary:"删除用户"`
	adminin.MemberDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新用户状态
type StatusReq struct {
	g.Meta `path:"/member/status" method:"post" tags:"用户" summary:"更新用户状态"`
	adminin.MemberStatusInp
}

type StatusRes struct{}

// SelectReq 获取可选的后台用户选项
type SelectReq struct {
	g.Meta `path:"/member/option" method:"get" tags:"用户" summary:"获取可选的后台用户选项"`
	adminin.MemberSelectInp
}

type SelectRes []*adminin.MemberSelectModel

// InfoReq 获取登录用户信息
type InfoReq struct {
	g.Meta `path:"/member/info" method:"get" tags:"用户" summary:"获取登录用户信息"`
}

type InfoRes struct {
	*adminin.LoginMemberInfoModel
}

// AddBalanceReq 增加余额
type AddBalanceReq struct {
	g.Meta `path:"/member/addBalance" method:"post" tags:"用户" summary:"增加余额"`
	adminin.MemberAddBalanceInp
}

type AddBalanceRes struct{}

// AddIntegralReq 增加积分
type AddIntegralReq struct {
	g.Meta `path:"/member/addIntegral" method:"post" tags:"用户" summary:"增加积分"`
	adminin.MemberAddIntegralInp
}

type AddIntegralRes struct{}
