// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMemberDao is the data access object for table hg_admin_member.
type AdminMemberDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns AdminMemberColumns // columns contains all the column names of Table for convenient usage.
}

// AdminMemberColumns defines and stores column names for table hg_admin_member.
type AdminMemberColumns struct {
	Id                 string // 管理员ID
	DeptId             string // 部门ID
	RoleId             string // 角色ID
	RealName           string // 真实姓名
	Username           string // 帐号
	PasswordHash       string // 密码
	Salt               string // 密码盐
	PasswordResetToken string // 密码重置令牌
	Integral           string // 积分
	Balance            string // 余额
	Avatar             string // 头像
	Sex                string // 性别
	Qq                 string // qq
	Email              string // 邮箱
	Mobile             string // 手机号码
	Birthday           string // 生日
	CityId             string // 城市编码
	Address            string // 联系地址
	Pid                string // 上级管理员ID
	Level              string // 关系树等级
	Tree               string // 关系树
	InviteCode         string // 邀请码
	Cash               string // 提现配置
	LastActiveAt       string // 最后活跃时间
	Remark             string // 备注
	Status             string // 状态
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
}

// adminMemberColumns holds the columns for table hg_admin_member.
var adminMemberColumns = AdminMemberColumns{
	Id:                 "id",
	DeptId:             "dept_id",
	RoleId:             "role_id",
	RealName:           "real_name",
	Username:           "username",
	PasswordHash:       "password_hash",
	Salt:               "salt",
	PasswordResetToken: "password_reset_token",
	Integral:           "integral",
	Balance:            "balance",
	Avatar:             "avatar",
	Sex:                "sex",
	Qq:                 "qq",
	Email:              "email",
	Mobile:             "mobile",
	Birthday:           "birthday",
	CityId:             "city_id",
	Address:            "address",
	Pid:                "pid",
	Level:              "level",
	Tree:               "tree",
	InviteCode:         "invite_code",
	Cash:               "cash",
	LastActiveAt:       "last_active_at",
	Remark:             "remark",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewAdminMemberDao creates and returns a new DAO object for table data access.
func NewAdminMemberDao() *AdminMemberDao {
	return &AdminMemberDao{
		group:   "default",
		table:   "hg_admin_member",
		columns: adminMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminMemberDao) Columns() AdminMemberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
