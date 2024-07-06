// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentDao is the data access object for table hg_sys_attachment.
type SysAttachmentDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysAttachmentColumns // columns contains all the column names of Table for convenient usage.
}

// SysAttachmentColumns defines and stores column names for table hg_sys_attachment.
type SysAttachmentColumns struct {
	Id        string // 文件ID
	AppId     string // 应用ID
	MemberId  string // 管理员ID
	CateId    string // 上传分类
	Drive     string // 上传驱动
	Name      string // 文件原始名
	Kind      string // 上传类型
	MimeType  string // 扩展类型
	NaiveType string // NaiveUI类型
	Path      string // 本地路径
	FileUrl   string // url
	Size      string // 文件大小
	Ext       string // 扩展名
	Md5       string // md5校验码
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// sysAttachmentColumns holds the columns for table hg_sys_attachment.
var sysAttachmentColumns = SysAttachmentColumns{
	Id:        "id",
	AppId:     "app_id",
	MemberId:  "member_id",
	CateId:    "cate_id",
	Drive:     "drive",
	Name:      "name",
	Kind:      "kind",
	MimeType:  "mime_type",
	NaiveType: "naive_type",
	Path:      "path",
	FileUrl:   "file_url",
	Size:      "size",
	Ext:       "ext",
	Md5:       "md5",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysAttachmentDao creates and returns a new DAO object for table data access.
func NewSysAttachmentDao() *SysAttachmentDao {
	return &SysAttachmentDao{
		group:   "default",
		table:   "hg_sys_attachment",
		columns: sysAttachmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAttachmentDao) Columns() SysAttachmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAttachmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
