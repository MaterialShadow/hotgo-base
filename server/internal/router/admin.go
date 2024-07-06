// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/controller/admin/admin"
	"hotgo/internal/controller/admin/common"
	"hotgo/internal/controller/admin/sys"
	"hotgo/internal/router/genrouter"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

func Admin(ctx context.Context, group *ghttp.RouterGroup) {
	// 兼容后台登录入口
	group.ALL("/login", func(r *ghttp.Request) {
		r.Response.RedirectTo("/admin")
	})

	group.Group(simple.RouterPrefix(ctx, consts.AppAdmin), func(group *ghttp.RouterGroup) {
		group.Bind(
			common.Site, // 基础
		)
		group.Middleware(service.Middleware().AdminAuth)
		group.Bind(
			common.Console, // 控制台
			common.Upload,  // 上传
			common.Wechat,  // 微信授权
			sys.Config,     // 配置
			sys.DictType,   // 字典类型
			sys.DictData,   // 字典数据
			sys.Attachment, // 附件
			admin.Member,   // 用户
			admin.Monitor,  // 监控
			admin.Role,     // 路由
			admin.Dept,     // 部门
			admin.Menu,     // 菜单
			admin.Post,     // 岗位
		)

	})

	// 注册生成路由
	genrouter.Register(ctx, group)
}
