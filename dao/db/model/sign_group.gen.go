// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
	"gorm.io/plugin/soft_delete"
)

const TableNameSignGroup = "sign_group"

// SignGroup mapped from table <sign_group>
type SignGroup struct {
	ID        int                   `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name      string                `gorm:"column:name;type:varchar(50);not null" json:"name"`
	CreatedAt *time.Time            `gorm:"column:created_at;type:int unsigned;autoCreateTime" json:"created_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	UpdatedAt *time.Time            `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	Start     time.Time             `gorm:"column:start;type:datetime;not null" json:"start"`
	End       time.Time             `gorm:"column:end;type:datetime;not null" json:"end"`
	Count     *int64                `gorm:"column:count;type:bigint unsigned;default:1" json:"count"`
	Avater    *string               `gorm:"column:avater;type:varchar(50)" json:"avater"`
	Owner     int64                 `gorm:"column:owner;type:bigint unsigned;not null" json:"owner"`
}

// TableName SignGroup's table name
func (*SignGroup) TableName() string {
	return TableNameSignGroup
}