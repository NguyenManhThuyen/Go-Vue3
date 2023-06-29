package routes

import (
	"api/modules/classes/controller"

	"github.com/gofiber/fiber/v2"
)

func InitClassRoutes(app *fiber.App) {
	class := app.Group("/users/classes")
	class.Get("/", controller.GetClass)
	class.Get("/:id", controller.GetClassByID)
	class.Post("/", controller.CreateClass)
	class.Put("/:id", controller.UpdateClass)
	class.Delete("/:id",controller.DeleteClassByID)
}
