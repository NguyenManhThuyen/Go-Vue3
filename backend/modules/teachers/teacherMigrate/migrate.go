package teacherMigrate

import (
	"api/database"
	model "api/modules/teachers/model"
)

func MigrateTeacher() bool {
	db := database.DB

	db.AutoMigrate(&model.Teacher{})

	return true
}