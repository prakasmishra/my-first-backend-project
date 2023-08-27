package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:password@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=true"

func InitGorm() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
}
