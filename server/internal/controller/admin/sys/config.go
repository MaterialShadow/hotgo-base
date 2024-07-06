// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/config"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Config = cConfig{}
)

type cConfig struct{}

// GetConfig 获取指定分组的配置
func (c *cConfig) GetConfig(ctx context.Context, req *config.GetReq) (res *config.GetRes, err error) {
	res = new(config.GetRes)
	res.GetConfigModel, err = service.SysConfig().GetConfigByGroup(ctx, &req.GetConfigInp)
	return
}

// UpdateConfig 更新指定分组的配置
func (c *cConfig) UpdateConfig(ctx context.Context, req *config.UpdateReq) (res *config.UpdateRes, err error) {
	err = service.SysConfig().UpdateConfigByGroup(ctx, &req.UpdateConfigInp)
	return
}

// TypeSelect 数据类型选项
func (c *cConfig) TypeSelect(_ context.Context, _ *config.TypeSelectReq) (res config.TypeSelectRes, err error) {
	for _, v := range consts.ConfigTypes {
		res = append(res, form.Select{
			Value: v,
			Name:  v,
			Label: v,
		})
	}
	return
}

// GetCash 获取提现的配置
func (c *cConfig) GetCash(ctx context.Context, _ *config.GetCashReq) (res *config.GetCashRes, err error) {
	res = new(config.GetCashRes)
	res.GetConfigModel, err = service.SysConfig().GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "cash"})
	return
}
