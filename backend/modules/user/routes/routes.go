package routes

import (
	"api/modules/user/controller"

	"github.com/gofiber/fiber/v2"
)

func InitTeacherRoutes(app *fiber.App) {
	user := app.Group("/users/user")
	user.Get("/", controller.GetUser)
	user.Get("/:id", controller.GetUserByID)
	user.Post("/", controller.CreateUser)
	user.Put("/:id", controller.UpdateUser)
	user.Delete("/:id",controller.DeleteUserByID)

}
