// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package social

const TableNameClubinvite = "clubinvite"

// Clubinvite 俱乐部邀请表
type Clubinvite struct {
	ID         int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:俱乐部邀请记录唯一标识" json:"id"` // 俱乐部邀请记录唯一标识
	ClubID     int64 `gorm:"column:club_id;comment:邀请加入的俱乐部ID" json:"club_id"`                      // 邀请加入的俱乐部ID
	SenderUID  int64 `gorm:"column:sender_uid;comment:邀请人用户ID" json:"sender_uid"`                   // 邀请人用户ID
	TargetUID  int64 `gorm:"column:target_uid;comment:被邀请人用户ID" json:"target_uid"`                  // 被邀请人用户ID
	InviteTime int64 `gorm:"column:invite_time;comment:邀请时间戳" json:"invite_time"`                   // 邀请时间戳
	Status     int32 `gorm:"column:status;comment:邀请状态（0: 待处理, 1: 已同意, 2: 已拒绝）" json:"status"`      // 邀请状态（0: 待处理, 1: 已同意, 2: 已拒绝）
}

// TableName Clubinvite's table name
func (*Clubinvite) TableName() string {
	return TableNameClubinvite
}
