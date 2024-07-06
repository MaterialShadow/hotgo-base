// Package contexts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package contexts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/model"
)

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
func Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextHTTPKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextHTTPKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func SetUser(ctx context.Context, user *model.Identity) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetUser, c == nil ")
		return
	}
	c.User = user
}

// SetResponse 设置组件响应 用于访问日志使用
func SetResponse(ctx context.Context, response *model.Response) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetResponse, c == nil ")
		return
	}
	c.Response = response
}

// SetModule 设置应用模块
func SetModule(ctx context.Context, module string) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetModule, c == nil ")
		return
	}
	c.Module = module
}

// GetUser 获取用户信息
func GetUser(ctx context.Context) *model.Identity {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.User
}

// GetUserId 获取用户ID
func GetUserId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.Id
}

// GetRoleId 获取用户角色ID
func GetRoleId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.RoleId
}

// GetRoleKey 获取用户角色唯一编码
func GetRoleKey(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.RoleKey
}

// GetDeptType 获取用户部门类型
func GetDeptType(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.DeptType
}

// IsCompanyDept 是否为公司部门
func IsCompanyDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeCompany
}

// IsTenantDept 是否为租户部门
func IsTenantDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeTenant
}

// IsMerchantDept 是否为商户部门
func IsMerchantDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeMerchant
}

// IsUserDept 是否为普通用户部门
func IsUserDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeUser
}

// GetModule 获取应用模块
func GetModule(ctx context.Context) string {
	c := Get(ctx)
	if c == nil {
		return ""
	}
	return c.Module
}

// SetAddonName 设置插件信息
func SetAddonName(ctx context.Context, name string) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetAddonName, c == nil ")
		return
	}
	Get(ctx).AddonName = name
}

// IsAddonRequest 是否为插件模块请求
func IsAddonRequest(ctx context.Context) bool {
	c := Get(ctx)
	if c == nil {
		return false
	}
	return GetAddonName(ctx) != ""
}

// GetAddonName 获取插件信息
func GetAddonName(ctx context.Context) string {
	c := Get(ctx)
	if c == nil {
		return ""
	}
	return Get(ctx).AddonName
}

// SetData 设置额外数据
func SetData(ctx context.Context, k string, v interface{}) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetData, c == nil ")
		return
	}
	Get(ctx).Data[k] = v
}

// SetDataMap 设置额外数据
func SetDataMap(ctx context.Context, vs g.Map) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetData, c == nil ")
		return
	}

	for k, v := range vs {
		Get(ctx).Data[k] = v
	}
}

// GetData 获取额外数据
func GetData(ctx context.Context) g.Map {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.Data
}
