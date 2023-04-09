// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
	"gorm.io/plugin/soft_delete"
)

const TableNameUserGroup = "user_group"

// UserGroup mapped from table <user_group>
type UserGroup struct {
	ID        int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt *time.Time            `gorm:"column:created_at;type:int unsigned;autoCreateTime" json:"created_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	UpdatedAt *time.Time            `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	UID       int64                 `gorm:"column:uid;type:bigint unsigned;not null" json:"uid"`
	Gid       int32                   `gorm:"column:gid;type:int unsigned;not null" json:"gid"`
	Score     *int64                `gorm:"column:score;type:bigint unsigned" json:"score"`
	Count     *int                  `gorm:"column:count;type:int unsigned" json:"count"`
}

// TableName UserGroup's table name
func (*UserGroup) TableName() string {
	return TableNameUserGroup
}