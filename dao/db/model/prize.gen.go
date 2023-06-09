// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
	"gorm.io/plugin/soft_delete"
)

const TableNamePrize = "prize"

// Prize mapped from table <prize>
type Prize struct {
	ID        int                   `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt *time.Time            `gorm:"column:created_at;type:int unsigned;autoCreateTime" json:"created_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	UpdatedAt *time.Time            `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	Name      string                `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Num       *int64                `gorm:"column:num;type:bigint unsigned" json:"num"`
	Picture   *string               `gorm:"column:picture;type:varchar(30)" json:"picture"`
	Aid       int                   `gorm:"column:aid;type:int unsigned;not null" json:"aid"`
}

// TableName Prize's table name
func (*Prize) TableName() string {
	return TableNamePrize
}
