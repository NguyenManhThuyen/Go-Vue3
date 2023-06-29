package modules

import (
	_ "api/docs"
	// authenRoute "api/modules/authen/routes"
	studentRoute "api/modules/student/routes"
	classRoute "api/modules/classes/routes"
	teacherRoute "api/modules/teachers/routes"
	tblRoute "api/modules/tbl/routes"
	userRoute "api/modules/user/routes"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	// authenRoute.InitAuthenRoutes(app)
	studentRoute.InitStudentsRoutes(app)
	classRoute.InitClassRoutes(app)
	teacherRoute.InitTeacherRoutes(app)
	tblRoute.InitTblRoutes(app)
	userRoute.InitTeacherRoutes(app)
}
