// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/utility/convert"
)

// MemberInfo 后台用户信息
var MemberInfo = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}
		for i, record := range result {
			// 部门
			if !record["dept_id"].IsEmpty() {
				deptName, err := g.Model("admin_dept").Ctx(ctx).
					Fields("name").
					Where("id", record["dept_id"]).
					Value()
				if err != nil {
					break
				}
				record["deptName"] = deptName
			}

			// 角色
			if !record["role_id"].IsEmpty() {
				roleName, err := g.Model("admin_role").Ctx(ctx).
					Fields("name").
					Where("id", record["role_id"]).
					Value()
				if err != nil {
					break
				}
				record["roleName"] = roleName
			}

			if !record["password_hash"].IsEmpty() {
				record["password_hash"] = gvar.New("")
			}

			if !record["salt"].IsEmpty() {
				record["salt"] = gvar.New("")
			}

			if !record["auth_key"].IsEmpty() {
				record["auth_key"] = gvar.New("")
			}

			result[i] = record
		}
		return
	},
}

type MemberSumma struct {
	Id       int64  `json:"id"                 description:"管理员ID"`
	RealName string `json:"realName"           description:"真实姓名"`
	Username string `json:"username"           description:"帐号"`
	Avatar   string `json:"avatar"             description:"头像"`
}

// MemberSummary 操作人摘要信息
var MemberSummary = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}

		var (
			createdByIds []int64
			updatedByIds []int64
			memberIds    []int64
		)

		for _, record := range result {
			if record["created_by"].Int64() > 0 {
				createdByIds = append(createdByIds, record["created_by"].Int64())
			}
			if record["updated_by"].Int64() > 0 {
				updatedByIds = append(updatedByIds, record["updated_by"].Int64())
			}
			if record["member_id"].Int64() > 0 {
				memberIds = append(memberIds, record["member_id"].Int64())
			}
		}

		memberIds = append(memberIds, createdByIds...)
		memberIds = append(memberIds, updatedByIds...)
		memberIds = convert.UniqueSlice(memberIds)
		if len(memberIds) == 0 {
			return
		}

		var members []*MemberSumma
		if err = g.Model("admin_member").Ctx(ctx).WhereIn("id", memberIds).Scan(&members); err != nil {
			return nil, err
		}

		if len(members) == 0 {
			return
		}

		findMember := func(id *gvar.Var) *MemberSumma {
			for _, v := range members {
				if v.Id == id.Int64() {
					return v
				}
			}
			return nil
		}

		for _, record := range result {
			if record["created_by"].Int64() > 0 {
				record["createdBySumma"] = gvar.New(findMember(record["created_by"]))
			}
			if record["updated_by"].Int64() > 0 {
				record["updatedBySumma"] = gvar.New(findMember(record["updated_by"]))
			}
			if record["member_id"].Int64() > 0 {
				record["memberBySumma"] = gvar.New(findMember(record["member_id"]))
			}
		}
		return
	},
}
