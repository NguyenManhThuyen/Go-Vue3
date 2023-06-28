package tblMigrate

import (
	"api/database"
	model "api/modules/tbl/model"
)

func MigrateTbl() bool {
	db := database.DB

	db.AutoMigrate(&model.Tbl{})

	return true
}