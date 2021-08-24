package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db = ConnectDB()
)

func ConnectDB() *gorm.DB {
	dsn := "root:duyngo99@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func CreateTables() {
	if (!db.Migrator().HasTable(&Order{})) {
		db.AutoMigrate(&Customer{}, &Order{}, &Product{}, &OrderDetail{})
	}
}
