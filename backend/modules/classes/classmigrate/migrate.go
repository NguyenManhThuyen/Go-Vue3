package classmigrate

import (
	"api/database"
	model "api/modules/classes/model"
)

func MigrateClass() bool {
	db := database.DB

	db.AutoMigrate(&model.Class{})

	return true
}
