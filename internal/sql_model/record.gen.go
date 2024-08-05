// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRecord = "record"

// Record mapped from table <record>
type Record struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:金币变化记录" json:"id"` // 金币变化记录
	UID       int64          `gorm:"column:uid;comment:充值者ID" json:"uid"`                              // 充值者ID
	Tid       int64          `gorm:"column:tid;comment:牌桌ID" json:"tid"`                               // 牌桌ID
	Pergold   int64          `gorm:"column:pergold;comment:支付之前" json:"pergold"`                       // 支付之前
	Payment   int64          `gorm:"column:payment;comment:支付" json:"payment"`                         // 支付
	Gold      int64          `gorm:"column:gold;comment:金币" json:"gold"`                               // 金币
	Code      int32          `gorm:"column:code;comment:操作码" json:"code"`                              // 操作码
	Order     string         `gorm:"column:order;comment:订单号(牌局号)" json:"order"`                       // 订单号(牌局号)
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
