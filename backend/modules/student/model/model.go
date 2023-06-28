package model

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID      int
	Name    string
	ClassID int
	Age     int
	Address string
}

func (Student) TableName() string {
	return "students"
}
