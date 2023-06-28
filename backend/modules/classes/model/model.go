package model

import (
	
)


type Class struct {
	ID           int
	Name         string
	TeacherID    int
	StudentCount int
}

func (Class) TableName() string {
	return "classes"
}
