package modules

import (
	"api/modules/authen/migrate"
	"api/modules/classes/classmigrate"
	"api/modules/student/studentmigrate"
	"api/modules/tbl/tblMigrate"
	"api/modules/teachers/teacherMigrate"
	"api/modules/user/userMigrate"
)

func MigrateModule() bool {
	migrate.MigrateAuthen()
	studentmigrate.MigrateStudent()
	classmigrate.MigrateClass()
	teacherMigrate.MigrateTeacher()
	tblMigrate.MigrateTbl()
	userMigrate.MigrateUser()
	return true
}
