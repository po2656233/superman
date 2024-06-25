// Code generated by gorm.io/genSqlModel. DO NOT EDIT.
// Code generated by gorm.io/genSqlModel. DO NOT EDIT.
// Code generated by gorm.io/genSqlModel. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePlatform = "platform"

// Platform mapped from table <platform>
type Platform struct {
	ID         int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"` // id
	Name       string         `gorm:"column:name;comment:名字" json:"name"`                           // 名字
	Usercount  int32          `gorm:"column:usercount;comment:在线人数" json:"usercount"`               // 在线人数
	Rooms      string         `gorm:"column:rooms;comment:房间编号以逗号分隔" json:"rooms"`                  // 房间编号以逗号分隔
	Activities string         `gorm:"column:activities;comment:日常活动" json:"activities"`             // 日常活动
	Remark     string         `gorm:"column:remark;comment:备注" json:"remark"`                       // 备注
	Code       string         `gorm:"column:code;comment:与用户表里的平台ID对应" json:"code"`                 // 与用户表里的平台ID对应
	Servers    string         `gorm:"column:servers;comment:服务器地址" json:"servers"`                  // 服务器地址
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy   int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy   int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Platform's table name
func (*Platform) TableName() string {
	return TableNamePlatform
}
