// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package social

const TableNameFriendapply = "friendapply"

// Friendapply 好友申请表
type Friendapply struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:好友申请记录唯一标识" json:"id"` // 好友申请记录唯一标识
	SenderUID int64  `gorm:"column:sender_uid;comment:申请人用户ID" json:"sender_uid"`                  // 申请人用户ID
	TargetUID int64  `gorm:"column:target_uid;comment:被申请人用户ID" json:"target_uid"`                 // 被申请人用户ID
	Cont      string `gorm:"column:cont;comment:申请附言" json:"cont"`                                 // 申请附言
	ApplyTime int64  `gorm:"column:apply_time;comment:申请时间戳" json:"apply_time"`                    // 申请时间戳
	Status    int32  `gorm:"column:status;comment:申请状态（0: 待处理, 1: 已同意, 2: 已拒绝）" json:"status"`     // 申请状态（0: 待处理, 1: 已同意, 2: 已拒绝）
}

// TableName Friendapply's table name
func (*Friendapply) TableName() string {
	return TableNameFriendapply
}
