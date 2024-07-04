// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRecord = "record"

// Record mapped from table <record>
type Record struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:金币变化记录" json:"id"` // 金币变化记录
	UID       int64          `gorm:"column:uid;comment:充值者ID" json:"uid"`                              // 充值者ID
	Byid      int64          `gorm:"column:byid;comment:给充值者充值的ID" json:"byid"`                        // 给充值者充值的ID
	HostID    int64          `gorm:"column:hostID;comment:房主ID" json:"hostID"`                         // 房主ID
	Gid       int64          `gorm:"column:gid;comment:游戏ID" json:"gid"`                               // 游戏ID
	Kid       int64          `gorm:"column:kid;comment:种类ID" json:"kid"`                               // 种类ID
	Level     int32          `gorm:"column:level;comment:等级" json:"level"`                             // 等级
	Pergold   int64          `gorm:"column:pergold;comment:支付之前" json:"pergold"`                       // 支付之前
	Payment   int64          `gorm:"column:payment;comment:支付" json:"payment"`                         // 支付
	Gold      int64          `gorm:"column:gold;comment:金币" json:"gold"`                               // 金币
	Time      int64          `gorm:"column:time;comment:时间戳" json:"time"`                              // 时间戳
	Code      int32          `gorm:"column:code;comment:操作码" json:"code"`                              // 操作码
	Rating    int32          `gorm:"column:rating;comment:评级|额定|差饷(父级)" json:"rating"`                 // 评级|额定|差饷(父级)
	Success   int32          `gorm:"column:success;comment:是否成功" json:"success"`                       // 是否成功
	Order     string         `gorm:"column:order;comment:订单号" json:"order"`                            // 订单号
	Remark    string         `gorm:"column:remark;comment:备注" json:"remark"`                           // 备注
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy  int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy  int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}
