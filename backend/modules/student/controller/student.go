package controller

import (
	"api/database"
	"api/modules/student/model"
	"github.com/gofiber/fiber/v2"
)

// GetStudent trả về danh sách tất cả học sinh
// @Summary Lấy danh sách học sinh
// @Description Trả về danh sách tất cả học sinh
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {array} model.Student
// @Failure 500 {object} fiber.Map
// @Router /users/students [get]
func GetStudent(c *fiber.Ctx) error {
	db := database.DB
	var students []model.Student
	if err := db.Find(&students).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve students",
		})
	}
	return c.JSON(students)
}

// GetStudentByID trả về thông tin của một học sinh dựa trên ID
// @Summary Lấy thông tin học sinh
// @Description Trả về thông tin của một học sinh dựa trên ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "ID của học sinh"
// @Success 200 {object} model.Student
// @Failure 404 {object} fiber.Map
// @Router /users/students/{id} [get]
func GetStudentByID(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student
	if err := db.First(&student, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}
	return c.JSON(student)
}

// CreateStudent tạo một học sinh mới
// @Summary Tạo học sinh mới
// @Description Tạo một học sinh mới
// @Tags Students
// @Accept json
// @Produce json
// @Param student body model.Student true "Thông tin học sinh mới"
// @Success 201 {object} model.Student
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/students [post]
func CreateStudent(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}
	if err := db.Create(&student).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create student",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(student)
}

// UpdateStudent cập nhật thông tin một học sinh dựa trên ID
// @Summary Cập nhật thông tin học sinh
// @Description Cập nhật thông tin một học sinh dựa trên ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "ID của học sinh"
// @Param student body model.Student true "Thông tin học sinh cần cập nhật"
// @Success 200 {object} model.Student
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/students/{id} [put]
func UpdateStudent(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student
	if err := db.First(&student, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}
	if err := db.Save(&student).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update student",
		})
	}
	return c.JSON(student)
}

// DeleteStudentByID xóa một học sinh dựa trên ID
// @Summary Xóa học sinh
// @Description Xóa một học sinh dựa trên ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "ID của học sinh"
// @Success 200 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /users/students/{id} [delete]
func DeleteStudentByID(c *fiber.Ctx) error {
	db := database.DB
	var student model.Student
	if err := db.First(&student, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}
	if err := db.Delete(&student).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete student",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Student deleted",
	})
}


