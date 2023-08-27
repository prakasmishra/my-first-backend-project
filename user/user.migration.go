package user

import (
	"goproject/db"
	"log"
)

func InitMigration() {
	err := db.DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Auto migration failed: ", err)
	}
}
