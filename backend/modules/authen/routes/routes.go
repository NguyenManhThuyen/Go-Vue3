package moduleauthen

import (
	_ "api/docs"
	"api/middleware"
	authenController "api/modules/authen/controller"

	"github.com/gofiber/fiber/v2"
)

func InitAuthenRoutes(app *fiber.App) {
	/**
	*
	*	System User
	* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	*
	**/
	api := app.Group("/", middleware.AppInfo)

	/**
	*
	*	Authen
	* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	*
	**/
	api.Post("/login", authenController.Login)
	api.Post("/check-token", middleware.AppAuthen, authenController.CheckToken)

}
