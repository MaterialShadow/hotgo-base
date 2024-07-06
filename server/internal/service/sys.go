// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
)

type (
	ISysAddonsConfig interface {
		// GetConfigByGroup 获取指定分组的配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetAddonsConfigInp) (res *sysin.GetAddonsConfigModel, err error)
		// ConversionType 转换类型
		ConversionType(ctx context.Context, models *entity.SysAddonsConfig) (value interface{}, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateAddonsConfigInp) (err error)
	}
	ISysAttachment interface {
		// Model ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Delete 删除附件
		Delete(ctx context.Context, in *sysin.AttachmentDeleteInp) (err error)
		// View 获取附件信息
		View(ctx context.Context, in *sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error)
		// List 获取附件列表
		List(ctx context.Context, in *sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error)
		// ClearKind 清空上传类型
		ClearKind(ctx context.Context, in *sysin.AttachmentClearKindInp) (err error)
	}

	ISysConfig interface {
		// InitConfig 初始化系统配置
		InitConfig(ctx context.Context)
		// LoadConfig 加载系统配置
		LoadConfig(ctx context.Context) (err error)
		// GetLogin 获取登录配置
		GetLogin(ctx context.Context) (conf *model.LoginConfig, err error)
		// GetUpload 获取上传配置
		GetUpload(ctx context.Context) (conf *model.UploadConfig, err error)
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetLoadTCP 获取本地tcp配置
		GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error)
		// GetLoadGenerate 获取本地生成配置
		GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error)
		// GetLoadToken 获取本地token配置
		GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error)
		// GetLoadLog 获取本地日志配置
		GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error)
		// GetLoadServeLog 获取本地服务日志配置
		GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error)
		// GetConfigByGroup 获取指定分组的配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// ConversionType 转换类型
		ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) (err error)
		// ClusterSync 集群同步
		ClusterSync(ctx context.Context, message *gredis.Message)
	}

	ISysCronGroup interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.CronGroupDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.CronGroupEditInp) (err error)
		// Status 更新状态
		Status(ctx context.Context, in *sysin.CronGroupStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.CronGroupMaxSortInp) (res *sysin.CronGroupMaxSortModel, err error)
		// View 获取指定信息
		View(ctx context.Context, in *sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int, err error)
		// Select 选项
		Select(ctx context.Context, in *sysin.CronGroupSelectInp) (res *sysin.CronGroupSelectModel, err error)
	}
	ISysCurdDemo interface {
		// Model CURD列表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取CURD列表列表
		List(ctx context.Context, in *sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error)
		// Export 导出CURD列表
		Export(ctx context.Context, in *sysin.CurdDemoListInp) (err error)
		// Edit 修改/新增CURD列表
		Edit(ctx context.Context, in *sysin.CurdDemoEditInp) (err error)
		// Delete 删除CURD列表
		Delete(ctx context.Context, in *sysin.CurdDemoDeleteInp) (err error)
		// MaxSort 获取CURD列表最大排序
		MaxSort(ctx context.Context, in *sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error)
		// View 获取CURD列表指定信息
		View(ctx context.Context, in *sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error)
		// Switch 更新CURD列表开关
		Switch(ctx context.Context, in *sysin.CurdDemoSwitchInp) (err error)
	}
	ISysDictData interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.DictDataDeleteInp) error
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.DictDataEditInp) (err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.DictDataListInp) (list []*sysin.DictDataListModel, totalCount int, err error)
		// GetId 获取指定类型的ID
		GetId(ctx context.Context, t string) (id int64, err error)
		// GetType 获取指定ID的类型标识
		GetType(ctx context.Context, id int64) (types string, err error)
		// GetTypes 获取指定ID的所有类型标识，包含下级
		GetTypes(ctx context.Context, id int64) (types []string, err error)
		// Select 获取列表
		Select(ctx context.Context, in *sysin.DataSelectInp) (list sysin.DataSelectModel, err error)
	}
	ISysDictType interface {
		// Tree 树
		Tree(ctx context.Context) (list []*sysin.DictTypeTree, err error)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.DictTypeDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.DictTypeEditInp) (err error)
		// TreeSelect 获取类型关系树选项
		TreeSelect(ctx context.Context, in *sysin.DictTreeSelectInp) (list []*sysin.DictTypeTree, err error)
		// BuiltinSelect 内置字典选项
		BuiltinSelect() (list []*sysin.DictTypeTree)
	}
	ISysEmsLog interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.EmsLogDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.EmsLogEditInp) (err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *sysin.EmsLogStatusInp) (err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.EmsLogViewInp) (res *sysin.EmsLogViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.EmsLogListInp) (list []*sysin.EmsLogListModel, totalCount int, err error)
		// Send 发送邮件
		Send(ctx context.Context, in *sysin.SendEmsInp) (err error)
		// GetTemplate 获取指定邮件模板
		GetTemplate(ctx context.Context, template string, config *model.EmailConfig) (val string, err error)
		// AllowSend 是否允许发送
		AllowSend(ctx context.Context, models *entity.SysEmsLog, config *model.EmailConfig) (err error)
		// NowDayIpSendCount 当天IP累计发送次数
		NowDayIpSendCount(ctx context.Context, event string) (count int, err error)
		// VerifyCode 效验验证码
		VerifyCode(ctx context.Context, in *sysin.VerifyEmsCodeInp) (err error)
	}
	ISysGenCodes interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.GenCodesDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *sysin.GenCodesStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.GenCodesMaxSortInp) (res *sysin.GenCodesMaxSortModel, err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error)
		// Selects 选项
		Selects(ctx context.Context, in *sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error)
		// TableSelect 表选项
		TableSelect(ctx context.Context, in *sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error)
		// ColumnSelect 表字段选项
		ColumnSelect(ctx context.Context, in *sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error)
		// ColumnList 表字段列表
		ColumnList(ctx context.Context, in *sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error)
		// Preview 生成预览
		Preview(ctx context.Context, in *sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error)
		// Build 提交生成
		Build(ctx context.Context, in *sysin.GenCodesBuildInp) (err error)
	}
	ISysLog interface {
		// Model 请求日志Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Export 导出
		Export(ctx context.Context, in *sysin.LogListInp) (err error)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, log entity.SysLog) (err error)
		// AutoLog 根据配置自动记录请求日志
		AutoLog(ctx context.Context) error
		// AnalysisLog 解析日志数据
		AnalysisLog(ctx context.Context) entity.SysLog
		// View 获取指定请求日志信息
		View(ctx context.Context, in *sysin.LogViewInp) (res *sysin.LogViewModel, err error)
		// Delete 删除请求日志
		Delete(ctx context.Context, in *sysin.LogDeleteInp) (err error)
		// List 请求日志列表
		List(ctx context.Context, in *sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error)
	}
	ISysLoginLog interface {
		// Model 登录日志Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取登录日志列表
		List(ctx context.Context, in *sysin.LoginLogListInp) (list []*sysin.LoginLogListModel, totalCount int, err error)
		// Export 导出登录日志
		Export(ctx context.Context, in *sysin.LoginLogListInp) (err error)
		// Delete 删除登录日志
		Delete(ctx context.Context, in *sysin.LoginLogDeleteInp) (err error)
		// Push 推送登录日志
		Push(ctx context.Context, in *sysin.LoginLogPushInp)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, models entity.SysLoginLog) (err error)
	}
	ISysNormalTreeDemo interface {
		// Model 普通树表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取普通树表列表
		List(ctx context.Context, in *sysin.NormalTreeDemoListInp) (list []*sysin.NormalTreeDemoListModel, totalCount int, err error)
		// Edit 修改/新增普通树表
		Edit(ctx context.Context, in *sysin.NormalTreeDemoEditInp) (err error)
		// Delete 删除普通树表
		Delete(ctx context.Context, in *sysin.NormalTreeDemoDeleteInp) (err error)
		// MaxSort 获取普通树表最大排序
		MaxSort(ctx context.Context, in *sysin.NormalTreeDemoMaxSortInp) (res *sysin.NormalTreeDemoMaxSortModel, err error)
		// View 获取普通树表指定信息
		View(ctx context.Context, in *sysin.NormalTreeDemoViewInp) (res *sysin.NormalTreeDemoViewModel, err error)
		// TreeOption 获取普通树表关系树选项
		TreeOption(ctx context.Context) (nodes []tree.Node, err error)
	}
	ISysOptionTreeDemo interface {
		// Model 选项树表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取选项树表列表
		List(ctx context.Context, in *sysin.OptionTreeDemoListInp) (list []*sysin.OptionTreeDemoListModel, totalCount int, err error)
		// Edit 修改/新增选项树表
		Edit(ctx context.Context, in *sysin.OptionTreeDemoEditInp) (err error)
		// Delete 删除选项树表
		Delete(ctx context.Context, in *sysin.OptionTreeDemoDeleteInp) (err error)
		// MaxSort 获取选项树表最大排序
		MaxSort(ctx context.Context, in *sysin.OptionTreeDemoMaxSortInp) (res *sysin.OptionTreeDemoMaxSortModel, err error)
		// View 获取选项树表指定信息
		View(ctx context.Context, in *sysin.OptionTreeDemoViewInp) (res *sysin.OptionTreeDemoViewModel, err error)
		// TreeOption 获取选项树表关系树选项
		TreeOption(ctx context.Context) (nodes []tree.Node, err error)
	}
	ISysProvinces interface {
		// Tree 关系树选项列表
		Tree(ctx context.Context) (list []*sysin.ProvincesTree, err error)
		// Delete 删除省市区数据
		Delete(ctx context.Context, in *sysin.ProvincesDeleteInp) (err error)
		// Edit 修改/新增省市区数据
		Edit(ctx context.Context, in *sysin.ProvincesEditInp) (err error)
		// Status 更新省市区状态
		Status(ctx context.Context, in *sysin.ProvincesStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.ProvincesMaxSortInp) (res *sysin.ProvincesMaxSortModel, err error)
		// View 获取省市区信息
		View(ctx context.Context, in *sysin.ProvincesViewInp) (res *sysin.ProvincesViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.ProvincesListInp) (list []*sysin.ProvincesListModel, totalCount int, err error)
		// ChildrenList 获取省市区下级列表
		ChildrenList(ctx context.Context, in *sysin.ProvincesChildrenListInp) (list []*sysin.ProvincesChildrenListModel, totalCount int, err error)
		// UniqueId 获取省市区下级列表
		UniqueId(ctx context.Context, in *sysin.ProvincesUniqueIdInp) (res *sysin.ProvincesUniqueIdModel, err error)
		// Select 省市区选项
		Select(ctx context.Context, in *sysin.ProvincesSelectInp) (res *sysin.ProvincesSelectModel, err error)
	}
	ISysServeLicense interface {
		// Model 服务许可证ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取服务许可证列表
		List(ctx context.Context, in *sysin.ServeLicenseListInp) (list []*sysin.ServeLicenseListModel, totalCount int, err error)
		// Export 导出服务许可证
		Export(ctx context.Context, in *sysin.ServeLicenseListInp) (err error)
		// Edit 修改/新增服务许可证
		Edit(ctx context.Context, in *sysin.ServeLicenseEditInp) (err error)
		// Delete 删除服务许可证
		Delete(ctx context.Context, in *sysin.ServeLicenseDeleteInp) (err error)
		// View 获取服务许可证指定信息
		View(ctx context.Context, in *sysin.ServeLicenseViewInp) (res *sysin.ServeLicenseViewModel, err error)
		// Status 更新服务许可证状态
		Status(ctx context.Context, in *sysin.ServeLicenseStatusInp) (err error)
		// AssignRouter 分配服务许可证路由
		AssignRouter(ctx context.Context, in *sysin.ServeLicenseAssignRouterInp) (err error)
	}
	ISysServeLog interface {
		// Model 服务日志Orm模型
		Model(ctx context.Context) *gdb.Model
		// List 获取服务日志列表
		List(ctx context.Context, in *sysin.ServeLogListInp) (list []*sysin.ServeLogListModel, totalCount int, err error)
		// Export 导出服务日志
		Export(ctx context.Context, in *sysin.ServeLogListInp) (err error)
		// Delete 删除服务日志
		Delete(ctx context.Context, in *sysin.ServeLogDeleteInp) (err error)
		// View 获取服务日志指定信息
		View(ctx context.Context, in *sysin.ServeLogViewInp) (res *sysin.ServeLogViewModel, err error)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, models entity.SysServeLog) (err error)
	}
	ISysSmsLog interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.SmsLogDeleteInp) (err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.SmsLogViewInp) (res *sysin.SmsLogViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.SmsLogListInp) (list []*sysin.SmsLogListModel, totalCount int, err error)
		// SendCode 发送验证码
		SendCode(ctx context.Context, in *sysin.SendCodeInp) (err error)
		// GetTemplate 获取指定短信模板
		GetTemplate(ctx context.Context, template string, config *model.SmsConfig) (val string, err error)
		// AllowSend 是否允许发送
		AllowSend(ctx context.Context, models *entity.SysSmsLog, config *model.SmsConfig) (err error)
		// NowDayIpSendCount 当天IP累计发送次数
		NowDayIpSendCount(ctx context.Context, event string) (count int, err error)
		// VerifyCode 效验验证码
		VerifyCode(ctx context.Context, in *sysin.VerifyCodeInp) (err error)
	}
	ISysTestCategory interface {
		// Model 测试分类ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取测试分类列表
		List(ctx context.Context, in *sysin.TestCategoryListInp) (list []*sysin.TestCategoryListModel, totalCount int, err error)
		// Edit 修改/新增测试分类
		Edit(ctx context.Context, in *sysin.TestCategoryEditInp) (err error)
		// Delete 删除测试分类
		Delete(ctx context.Context, in *sysin.TestCategoryDeleteInp) (err error)
		// MaxSort 获取测试分类最大排序
		MaxSort(ctx context.Context, in *sysin.TestCategoryMaxSortInp) (res *sysin.TestCategoryMaxSortModel, err error)
		// View 获取测试分类指定信息
		View(ctx context.Context, in *sysin.TestCategoryViewInp) (res *sysin.TestCategoryViewModel, err error)
		// Status 更新测试分类状态
		Status(ctx context.Context, in *sysin.TestCategoryStatusInp) (err error)
		// Option 获取测试分类选项
		Option(ctx context.Context) (opts []*model.Option, err error)
	}
)

var (
	localSysAddonsConfig   ISysAddonsConfig
	localSysAttachment     ISysAttachment
	localSysConfig         ISysConfig
	localSysCronGroup      ISysCronGroup
	localSysCurdDemo       ISysCurdDemo
	localSysDictData       ISysDictData
	localSysDictType       ISysDictType
	localSysEmsLog         ISysEmsLog
	localSysGenCodes       ISysGenCodes
	localSysLog            ISysLog
	localSysLoginLog       ISysLoginLog
	localSysNormalTreeDemo ISysNormalTreeDemo
	localSysOptionTreeDemo ISysOptionTreeDemo
	localSysProvinces      ISysProvinces
	localSysServeLicense   ISysServeLicense
	localSysServeLog       ISysServeLog
	localSysSmsLog         ISysSmsLog
	localSysTestCategory   ISysTestCategory
)

func SysAddonsConfig() ISysAddonsConfig {
	if localSysAddonsConfig == nil {
		panic("implement not found for interface ISysAddonsConfig, forgot register?")
	}
	return localSysAddonsConfig
}

func RegisterSysAddonsConfig(i ISysAddonsConfig) {
	localSysAddonsConfig = i
}

func SysAttachment() ISysAttachment {
	if localSysAttachment == nil {
		panic("implement not found for interface ISysAttachment, forgot register?")
	}
	return localSysAttachment
}

func RegisterSysAttachment(i ISysAttachment) {
	localSysAttachment = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}


func SysCronGroup() ISysCronGroup {
	if localSysCronGroup == nil {
		panic("implement not found for interface ISysCronGroup, forgot register?")
	}
	return localSysCronGroup
}

func RegisterSysCronGroup(i ISysCronGroup) {
	localSysCronGroup = i
}

func SysCurdDemo() ISysCurdDemo {
	if localSysCurdDemo == nil {
		panic("implement not found for interface ISysCurdDemo, forgot register?")
	}
	return localSysCurdDemo
}

func RegisterSysCurdDemo(i ISysCurdDemo) {
	localSysCurdDemo = i
}

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}


