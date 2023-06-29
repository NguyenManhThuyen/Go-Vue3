package routes

import (
	"api/modules/student/controller"

	"github.com/gofiber/fiber/v2"
)

func InitStudentsRoutes(app *fiber.App) {
	student := app.Group("/users/students")
	student.Get("/", controller.GetStudent)
	student.Get("/:id", controller.GetStudentByID)
	student.Post("/", controller.CreateStudent)
	student.Put("/:id", controller.UpdateStudent)
	student.Delete("/:id",controller.DeleteStudentByID)
	// student.Get("/",controller.Test)
}
