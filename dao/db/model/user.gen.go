// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
	"gorm.io/plugin/soft_delete"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt *time.Time            `gorm:"column:created_at;type:int unsigned;autoCreateTime" json:"created_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	UpdatedAt *time.Time            `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	Email     string                `gorm:"column:email;type:varchar(50);not null" json:"email"`
	Name      string                `gorm:"column:name;type:varchar(30);not null" json:"name"`
	Password  string                `gorm:"column:password;type:varchar(100);not null" json:"password"`
	Avater    *string               `gorm:"column:avater;type:varchar(50)" json:"avater"`
	Address   *string               `gorm:"column:address;type:varchar(50)" json:"address"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
