package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Name string
	Age  int
}

func (User) TableName() string {
	return "test_user"
}
