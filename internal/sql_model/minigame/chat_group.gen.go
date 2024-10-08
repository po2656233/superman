// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package minigame

import (
	"time"

	"gorm.io/gorm"
)

const TableNameChatGroup = "chat_group"

// ChatGroup mapped from table <chat_group>
type ChatGroup struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string         `gorm:"column:name;comment:群名称" json:"name"`                                             // 群名称
	Hostid       int64          `gorm:"column:hostid;comment:群主ID" json:"hostid"`                                        // 群主ID
	Setuptime    int64          `gorm:"column:setuptime;comment:创建时间" json:"setuptime"`                                  // 创建时间
	Explain      string         `gorm:"column:explain;comment:群简介" json:"explain"`                                       // 群简介
	Notice       string         `gorm:"column:notice;comment:群公告" json:"notice"`                                         // 群公告
	Adminlist    string         `gorm:"column:adminlist;comment:管理者者" json:"adminlist"`                                  // 管理者者
	Memberlist   string         `gorm:"column:memberlist;comment:成员列表" json:"memberlist"`                                // 成员列表
	Applylist    string         `gorm:"column:applylist;comment:申请列表" json:"applylist"`                                  // 申请列表
	Bannedlist   string         `gorm:"column:bannedlist;comment:禁言者" json:"bannedlist"`                                 // 禁言者
	Robotid      int64          `gorm:"column:robotid;comment:机器人ID" json:"robotid"`                                     // 机器人ID
	Robotkey     string         `gorm:"column:robotkey;comment:机器人密钥" json:"robotkey"`                                   // 机器人密钥
	Robotcontrol int32          `gorm:"column:robotcontrol;comment:机器人控制(0:关闭 1:开启 2:停止 3:维护 4:销毁)" json:"robotcontrol"` // 机器人控制(0:关闭 1:开启 2:停止 3:维护 4:销毁)
	Remark       string         `gorm:"column:remark;comment:缓存key" json:"remark"`                                       // 缓存key
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy     int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy     int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName ChatGroup's table name
func (*ChatGroup) TableName() string {
	return TableNameChatGroup
}
