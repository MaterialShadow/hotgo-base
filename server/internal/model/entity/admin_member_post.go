// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdminMemberPost is the golang structure for table admin_member_post.
type AdminMemberPost struct {
	MemberId int64 `json:"memberId" orm:"member_id" description:"管理员ID"`
	PostId   int64 `json:"postId"   orm:"post_id"   description:"岗位ID"`
}
