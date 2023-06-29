package controller

import (
	"api/database"
	"api/modules/classes/model"

	"github.com/gofiber/fiber/v2"
)

// GetClass trả về danh sách tất cả Lớp học
// @Summary Lấy danh sách lớp học
// @Description Trả về danh sách tất cả lớp học
// @Tags Classes
// @Accept json
// @Produce json
// @Success 200 {array} model.Class
// @Failure 500 {object} fiber.Map
// @Router /users/classes [get]
func GetClass(c *fiber.Ctx) error {
	db := database.DB
	var classes []model.Class
	if err := db.Find(&classes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve classes",
		})
	}

	return c.JSON(classes)
}

// GetStudentByID trả về thông tin của một Lớp học dựa trên ID
// @Summary Lấy thông tin Lớp học
// @Description Trả về thông tin của một Lớp học dựa trên ID
// @Tags Classes
// @Accept json
// @Produce json
// @Param id path int true "ID của Lớp học"
// @Success 200 {object} model.Class
// @Failure 404 {object} fiber.Map
// @Router /users/classes/{id} [get]
func GetClassByID(c *fiber.Ctx) error {
	db := database.DB
	var class model.Class

	if err := db.First(&class, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Class not found",
		})
	}

	return c.JSON(class)
}

// CreateStudent tạo một Lớp học mới
// @Summary Tạo Lớp học mới
// @Description Tạo một Lớp học mới
// @Tags Classes
// @Accept json
// @Produce json
// @Param student body model.Class true "Thông tin Lớp học mới"
// @Success 201 {object} model.Class
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/classes [post]
func CreateClass(c *fiber.Ctx) error {
	db := database.DB
	var class model.Class
	if err := c.BodyParser(&class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	if err := db.Create(&class).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create class",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(class)
}

// UpdateStudent cập nhật thông tin một Lớp học dựa trên ID
// @Summary Cập nhật thông tin Lớp học
// @Description Cập nhật thông tin một Lớp học dựa trên ID
// @Tags Classes
// @Accept json
// @Produce json
// @Param id path int true "ID của Lớp học"
// @Param student body model.Class true "Thông tin Lớp học cần cập nhật"
// @Success 200 {object} model.Class
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/classes/{id} [put]
func UpdateClass(c *fiber.Ctx) error {
	db := database.DB
	var class model.Class
	if err := db.First(&class, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Class not found",
		})
	}

	if err := c.BodyParser(&class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	if err := db.Save(&class).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update class",
		})
	}

	return c.JSON(class)
}

// DeleteStudentByID xóa một Lớp học dựa trên ID
// @Summary Xóa Lớp học
// @Description Xóa một Lớp học dựa trên ID
// @Tags Classes
// @Accept json
// @Produce json
// @Param id path int true "ID của Lớp học"
// @Success 200 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /users/classes/{id} [delete]
func DeleteClassByID(c *fiber.Ctx) error {
	db := database.DB
	var class model.Class
	if err := db.First(&class, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Class not found",
		})
	}

	if err := db.Delete(&class).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete class",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Class deleted",
	})
}
