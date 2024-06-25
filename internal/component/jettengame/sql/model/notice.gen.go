// Code generated by gorm.io/genSqlModel. DO NOT EDIT.
// Code generated by gorm.io/genSqlModel. DO NOT EDIT.
// Code generated by gorm.io/genSqlModel. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNotice = "notice"

// Notice mapped from table <notice>
type Notice struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:通知ID" json:"id"` // 通知ID
	Type      int32          `gorm:"column:type;comment:通知类型" json:"type"`                           // 通知类型
	Platid    int64          `gorm:"column:platid;comment:平台ID" json:"platid"`                       // 平台ID
	Gameids   string         `gorm:"column:gameids;comment:指定游戏" json:"gameids"`                     // 指定游戏
	Kindid    int64          `gorm:"column:kindid;comment:关联游戏" json:"kindid"`                       // 关联游戏
	Level     int32          `gorm:"column:level;comment:游戏级别" json:"level"`                         // 游戏级别
	Title     string         `gorm:"column:title;comment:标题" json:"title"`                           // 标题
	Content   string         `gorm:"column:content;comment:内容" json:"content"`                       // 内容
	Start     int64          `gorm:"column:start;comment:起始时间" json:"start"`                         // 起始时间
	End       int64          `gorm:"column:end;comment:结束时间" json:"end"`                             // 结束时间
	Remark    string         `gorm:"column:remark;comment:备注" json:"remark"`                         // 备注
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy  int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy  int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Notice's table name
func (*Notice) TableName() string {
	return TableNameNotice
}
