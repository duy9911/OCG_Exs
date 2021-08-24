package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (db *gorm.DB, err error) {
	dns := "root:duyngo99@tcp(127.0.0.1:3306)/Class?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		fmt.Println(err)
		panic("Failed to open to the database!")
	}
	if !db.Migrator().HasTable(&Students) {
		db.Migrator().CreateTable(&Students)
	}
	return

}
