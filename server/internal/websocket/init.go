// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package websocket

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	mctx          = gctx.GetInitCtx()             // 上下文
	clientManager = NewClientManager()            // 客户端管理
	routers       = make(map[string]EventHandler) // 消息路由
	msgGo         = grpool.New(20)                // 消息处理协程池
)

// Start 启动
func Start() {
	go clientManager.start()
	go clientManager.ping()
	g.Log().Debug(mctx, "start websocket..")
}

// Stop 关闭
func Stop() {
	clientManager.closeSignal <- struct{}{}
}

// WsPage ws入口
func WsPage(r *ghttp.Request) {
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		return
	}
	currentTime := uint64(gtime.Now().Unix())
	client := NewClient(r, conn, currentTime)
	go client.read()
	go client.write()
	// 用户连接事件
	clientManager.Register <- client
}
