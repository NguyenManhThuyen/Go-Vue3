package routes

import (
	"api/controller"
	_ "api/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	// Document
	app.Get("/document/*", swagger.HandlerDefault)
	app.Get("/test", controller.Test)
}
