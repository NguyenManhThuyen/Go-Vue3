package controller

import (
	"api/database"
	"api/modules/tbl/model"


	"github.com/gofiber/fiber/v2"
)

func GetTbl(c *fiber.Ctx) error {
	db := database.DB
	var tbl []model.Tbl
		if err := db.Find(&tbl).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve tbl",
			})
		}

		return c.JSON(tbl)
}

func GetTblByID(c *fiber.Ctx) error {
	db := database.DB
	var tbl model.Tbl
		if err := db.First(&tbl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Teacher not found",
			})
		}

		return c.JSON(tbl)
}


func CreateTbl(c *fiber.Ctx) error {
	db := database.DB
	var tbl model.Tbl
		if err := c.BodyParser(&tbl); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		if err := db.Create(&tbl).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create teacher",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(tbl)
}


func UpdateTbl(c *fiber.Ctx) error {
	db := database.DB
	var tbl model.Tbl
		if err := db.First(&tbl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "tbl_department not found",
			})
		}

		if err := c.BodyParser(&tbl); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		if err := db.Save(&tbl).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update tbl_department",
			})
		}

		return c.JSON(tbl)
}


func DeleteTblByID(c *fiber.Ctx) error {
	db := database.DB
	var tbl model.Tbl
		if err := db.First(&tbl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "tbl_department not found",
			})
		}

		if err := db.Delete(&tbl).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete tbl_department",
			})
		}

		return c.JSON(fiber.Map{
			"message": "tbl_department deleted",
		})
}

