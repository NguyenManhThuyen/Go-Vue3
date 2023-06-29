package migrate

import (
	"api/database"
	model "api/modules/authen/model"
)

func MigrateAuthen() bool {
	db := database.DB

	db.AutoMigrate(&model.User{})

	return true
}