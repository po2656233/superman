// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTable = "table"

// Table mapped from table <table>
type Table struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:桌子ID" json:"id"`                // 桌子ID
	Name       string         `gorm:"column:name;comment:名称" json:"name"`                                            // 名称
	Rid        int64          `gorm:"column:rid;not null;comment:房间ID" json:"rid"`                                   // 房间ID
	Gid        int64          `gorm:"column:gid;not null;comment:游戏ID" json:"gid"`                                   // 游戏ID
	Opentime   int64          `gorm:"column:opentime;comment:开桌时间(时间戳)" json:"opentime"`                             // 开桌时间(时间戳)
	Taxation   int64          `gorm:"column:taxation;comment:固定收税" json:"taxation"`                                  // 固定收税
	Maxround   int32          `gorm:"column:maxround;comment:游戏最大轮次(=-1不受限)" json:"maxround"`                        // 游戏最大轮次(=-1不受限)
	Commission int32          `gorm:"column:commission;comment:台费(单位:万分之一,每局都会收取)" json:"commission"`                // 台费(单位:万分之一,每局都会收取)
	Remain     int32          `gorm:"column:remain;comment:剩余场次(=-1不受限)" json:"remain"`                              // 剩余场次(=-1不受限)
	Playscore  int64          `gorm:"column:playscore;comment:初始积分(携带的积分)" json:"playscore"`                         // 初始积分(携带的积分)
	MaxSitter  int32          `gorm:"column:max_sitter;default:-1;comment:客人(即:可容纳玩家数量 =-1时,不受限)" json:"max_sitter"` // 客人(即:可容纳玩家数量 =-1时,不受限)
	Remark     string         `gorm:"column:remark;comment:备注" json:"remark"`                                        // 备注
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy   int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy   int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Table's table name
func (*Table) TableName() string {
	return TableNameTable
}
