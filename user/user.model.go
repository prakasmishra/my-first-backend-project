package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FisrtName string
	LastName  string
	Email     string
}
