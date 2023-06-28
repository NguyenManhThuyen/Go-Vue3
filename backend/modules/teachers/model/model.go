package model

import (
)

type Teacher struct {
	ID      int
	Name    string
	Subject string
}


func (Teacher) TableName() string {
	return "teachers"
}
