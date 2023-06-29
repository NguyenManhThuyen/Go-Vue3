package studentmigrate

import (
	"api/database"
	model "api/modules/student/model"
)

func MigrateStudent() bool {
	db := database.DB

	db.AutoMigrate(&model.Student{})

	return true
}
