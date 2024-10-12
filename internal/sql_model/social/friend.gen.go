// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package social

const TableNameFriend = "friend"

// Friend 好友关系表
type Friend struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:好友关系记录唯一标识" json:"id"` // 好友关系记录唯一标识
	UID       int64 `gorm:"column:uid;comment:用户ID" json:"uid"`                                   // 用户ID
	FriendUID int64 `gorm:"column:friend_uid;comment:好友用户ID" json:"friend_uid"`                   // 好友用户ID
	IsBlack   int32 `gorm:"column:is_black;comment:是否加入黑名单" json:"is_black"`                      // 是否加入黑名单
	AddTime   int64 `gorm:"column:add_time;comment:添加好友的时间戳" json:"add_time"`                     // 添加好友的时间戳
	DelTime   int64 `gorm:"column:del_time;comment:删除好友的时间戳" json:"del_time"`                     // 删除好友的时间戳
}

// TableName Friend's table name
func (*Friend) TableName() string {
	return TableNameFriend
}
