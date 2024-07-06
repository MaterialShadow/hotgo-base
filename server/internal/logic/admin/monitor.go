// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/net"
	"hotgo/internal/library/location"
	"hotgo/internal/model"
	"hotgo/internal/service"
	"hotgo/utility/format"
	"hotgo/utility/simple"
	"runtime"
	"sync"
)

type sAdminMonitor struct {
	data *model.MonitorData
	sync.RWMutex
}

func NewAdminMonitor() *sAdminMonitor {
	return &sAdminMonitor{
		data: new(model.MonitorData),
	}
}

func init() {
	service.RegisterAdminMonitor(NewAdminMonitor())
}

// StartMonitor 启动服务监控
func (s *sAdminMonitor) StartMonitor(ctx context.Context) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		s.data.STartTime = gtime.Now().Timestamp()
		s.data.IntranetIP, _ = location.GetLocalIP()
		s.data.PublicIP, _ = location.GetPublicIP(ctx)

		_, err := gcron.Add(ctx, "@every 1s", func(ctx context.Context) {
			s.Lock()
			defer s.Unlock()
			s.netIO()
			s.loadAvg()
		}, "AdminMonitorCronJob")
		if err != nil {
			g.Log().Warningf(ctx, "StartMonitor CronJob err:%+v", err)
		}
	})
}

// GetMeta 获取监控元数据
func (s *sAdminMonitor) GetMeta(ctx context.Context) *model.MonitorData {
	return s.data
}

func (s *sAdminMonitor) loadAvg() {
	pl, _ := load.Avg()
	counter := model.LoadAvgStats{
		Time:  gtime.Now(),
		Avg:   pl.Load1,
		Ratio: pl.Load1 / (float64(runtime.NumCPU()) * 2) * 100,
	}

	s.data.LoadAvg = append(s.data.LoadAvg, &counter)
	if len(s.data.LoadAvg) > 10 {
		s.data.LoadAvg = append(s.data.LoadAvg[:0], s.data.LoadAvg[(1):]...)
	}
}

func (s *sAdminMonitor) netIO() {
	var counter model.NetIOCounters
	ni, _ := net.IOCounters(true)
	counter.Time = gtime.Now()
	for _, v := range ni {
		counter.BytesSent += v.BytesSent
		counter.BytesRecv += v.BytesRecv
	}

	if len(s.data.NetIO) > 0 {
		lastNetIO := s.data.NetIO[len(s.data.NetIO)-1]
		sub := counter.Time.Sub(lastNetIO.Time).Seconds()
		counter.Down = format.Round2Float64((float64(counter.BytesRecv - lastNetIO.BytesRecv)) / sub)
		counter.UP = format.Round2Float64((float64(counter.BytesSent - lastNetIO.BytesSent)) / sub)
	}

	s.data.NetIO = append(s.data.NetIO, &counter)
	if len(s.data.NetIO) > 10 {
		s.data.NetIO = append(s.data.NetIO[:0], s.data.NetIO[(1):]...)
	}
}
