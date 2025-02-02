// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/api/admin/monitor"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/internal/websocket"
	"hotgo/utility/simple"
	"hotgo/utility/useragent"
	"sort"
)

// Monitor 监控
var Monitor = cMonitor{
	wsManager: websocket.Manager(),
}

type cMonitor struct {
	wsManager *websocket.ClientManager
}

// UserOffline 下线用户
func (c *cMonitor) UserOffline(ctx context.Context, req *monitor.UserOfflineReq) (res *monitor.UserOfflineRes, err error) {
	client := c.wsManager.GetClient(req.Id)
	if client == nil {
		err = gerror.New("客户端已离线")
		return
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		websocket.SendSuccess(client, "kick")
		websocket.Close(client)
	})
	return
}

// UserOnlineList 获取用户在线列表
func (c *cMonitor) UserOnlineList(ctx context.Context, req *monitor.UserOnlineListReq) (res *monitor.UserOnlineListRes, err error) {
	var (
		clients []*monitor.UserOnlineModel
		i       int
	)

	if c.wsManager.GetClientsLen() == 0 {
		return
	}

	for conn := range c.wsManager.GetClients() {
		if conn.SendClose || conn.User == nil {
			continue
		}

		if req.UserId > 0 && req.UserId != conn.User.Id {
			continue
		}

		if req.Username != "" && req.Username != conn.User.Username {
			continue
		}

		if req.IP != "" && !gstr.Contains(conn.IP, req.IP) {
			continue
		}

		if len(req.FirstTime) == 2 && (conn.User.LoginAt.Before(req.FirstTime[0]) || conn.User.LoginAt.After(req.FirstTime[1])) {
			continue
		}

		clients = append(clients, &monitor.UserOnlineModel{
			ID:            conn.ID,
			IP:            conn.IP,
			Os:            useragent.GetOs(conn.UserAgent),
			Browser:       useragent.GetBrowser(conn.UserAgent),
			FirstTime:     conn.User.LoginAt.Unix(),
			HeartbeatTime: conn.HeartbeatTime,
			App:           conn.User.App,
			UserId:        conn.User.Id,
			Username:      conn.User.Username,
			Avatar:        conn.User.Avatar,
		})
	}

	res = new(monitor.UserOnlineListRes)
	res.PageRes.Pack(req, len(clients))

	sort.Slice(clients, func(i, j int) bool {
		if clients[i].FirstTime == clients[j].FirstTime {
			return clients[i].ID < clients[j].ID
		}
		return clients[i].FirstTime < clients[j].FirstTime
	})

	_, perPage, offset := form.CalPage(req.Page, req.PerPage)
	for k, v := range clients {
		if k >= offset && i <= perPage {
			if simple.IsDemo(ctx) {
				v.IP = consts.DemoTips
			}
			res.List = append(res.List, v)
			i++
		}
	}
	return
}
