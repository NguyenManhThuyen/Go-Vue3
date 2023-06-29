package routes

import (
	"api/modules/tbl/controller"

	"github.com/gofiber/fiber/v2"
)

func InitTblRoutes(app *fiber.App) {
	tbl := app.Group("/users/tbl")
	tbl.Get("/", controller.GetTbl)
	tbl.Get("/:id", controller.GetTblByID)
	tbl.Post("/", controller.CreateTbl)
	tbl.Put("/:id", controller.UpdateTbl)
	tbl.Delete("/:id",controller.DeleteTblByID)

}
