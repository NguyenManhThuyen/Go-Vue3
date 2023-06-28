package routes

import (
	"api/modules/teachers/controller"

	"github.com/gofiber/fiber/v2"
)

func InitTeacherRoutes(app *fiber.App) {
	teacher := app.Group("/users/teachers")
	teacher.Get("/", controller.GetTeacher)
	teacher.Get("/:id", controller.GetTeacherByID)
	teacher.Post("/", controller.CreateTeacher)
	teacher.Put("/:id", controller.UpdateTeacher)
	teacher.Delete("/:id",controller.DeleteTeacherByID)

}
