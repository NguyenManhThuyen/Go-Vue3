package userMigrate

import (
	"api/database"
	model "api/modules/user/model"
)

func MigrateUser() bool {
	db := database.DB

	db.AutoMigrate(&model.User{})

	return true
}