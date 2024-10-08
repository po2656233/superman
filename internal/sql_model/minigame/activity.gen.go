// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package minigame

import (
	"time"

	"gorm.io/gorm"
)

const TableNameActivity = "activity"

// Activity mapped from table <activity>
type Activity struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"` // id
	Name           string         `gorm:"column:name;comment:名字" json:"name"`                           // 名字
	Announcement   string         `gorm:"column:announcement;comment:公告" json:"announcement"`           // 公告
	StartTime      time.Time      `gorm:"column:startTime;comment:起始时间" json:"startTime"`               // 起始时间
	EndTime        time.Time      `gorm:"column:endTime;comment:结束时间" json:"endTime"`                   // 结束时间
	Intentionality string         `gorm:"column:intentionality;comment:目的" json:"intentionality"`       // 目的
	Prize          string         `gorm:"column:prize;comment:pd字节格式物品奖励" json:"prize"`                 // pd字节格式物品奖励
	Remark         string         `gorm:"column:remark;comment:备注" json:"remark"`                       // 备注
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy       int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy       int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Activity's table name
func (*Activity) TableName() string {
	return TableNameActivity
}
