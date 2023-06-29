package controller

import (
	"api/database"
	"api/modules/user/model"


	"github.com/gofiber/fiber/v2"
)


func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user []model.User
		if err := db.Find(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve user",
			})
		}

		return c.JSON(user)
}


func GetUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
		if err := db.First(&user, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
}


func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
		if err := db.First(&user, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update user",
			})
		}

		return c.JSON(user)
}


func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
		if err := db.First(&user, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		if err := db.Delete(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete user",
			})
		}

		return c.JSON(fiber.Map{
			"message": "user deleted",
		})
}

