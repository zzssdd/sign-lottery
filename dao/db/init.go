package db

import (
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Dao struct {
	User     User
	Admin    Admin
	Group    Group
	Sign     Sign
	Activity Activity
	Prize    Prize
	Order    Order
	Record   Record
}

func InitDB() {
	var err error
	db, err = gorm.Open(mysql.Open(constants.MysqlDSN), &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		Log.Errorln("open database err:", err)
	}
	mydb, _ := db.DB()
	mydb.SetMaxOpenConns(100)
	mydb.SetMaxIdleConns(10)
	mydb.SetConnMaxLifetime(10 * time.Second)
}

func NewDao() *Dao {
	if db == nil {
		InitDB()
	}
	return &Dao{
		User:     User{},
		Admin:    Admin{},
		Group:    Group{},
		Sign:     Sign{},
		Activity: Activity{},
		Prize:    Prize{},
		Order:    Order{},
		Record:   Record{},
	}
}
