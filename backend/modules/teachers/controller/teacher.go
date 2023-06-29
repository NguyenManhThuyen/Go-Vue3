package controller

import (
	"api/database"
	"api/modules/teachers/model"


	"github.com/gofiber/fiber/v2"
)

// GetStudent trả về danh sách tất cả giáo viên
// @Summary Lấy danh sách giáo viên
// @Description Trả về danh sách tất cả giáo viên
// @Tags Teachers
// @Accept json
// @Produce json
// @Success 200 {array} model.Teacher
// @Failure 500 {object} fiber.Map
// @Router /users/teachers [get]
func GetTeacher(c *fiber.Ctx) error {
	db := database.DB
	var teachers []model.Teacher
		if err := db.Find(&teachers).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve teachers",
			})
		}

		return c.JSON(teachers)
}

// GetStudentByID trả về thông tin của một giáo viên dựa trên ID
// @Summary Lấy thông tin giáo viên
// @Description Trả về thông tin của một giáo viên dựa trên ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int true "ID của giáo viên"
// @Success 200 {object} model.Teacher
// @Failure 404 {object} fiber.Map
// @Router /users/teachers/{id} [get]
func GetTeacherByID(c *fiber.Ctx) error {
	db := database.DB
	var teacher model.Teacher
		if err := db.First(&teacher, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Teacher not found",
			})
		}

		return c.JSON(teacher)
}

// CreateStudent tạo một giáo viên mới
// @Summary Tạo giáo viên mới
// @Description Tạo một giáo viên mới
// @Tags Teachers
// @Accept json
// @Produce json
// @Param teacher body model.Teacher true "Thông tin giáo viên mới"
// @Success 201 {object} model.Teacher
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/teachers [post]
func CreateTeacher(c *fiber.Ctx) error {
	db := database.DB
	var teacher model.Teacher
		if err := c.BodyParser(&teacher); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		if err := db.Create(&teacher).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create teacher",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(teacher)
}

// UpdateStudent cập nhật thông tin một giáo viên dựa trên ID
// @Summary Cập nhật thông tin giáo viên
// @Description Cập nhật thông tin một giáo viên dựa trên ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int true "ID của giáo viên"
// @Param teacher body model.Teacher true "Thông tin giáo viên cần cập nhật"
// @Success 200 {object} model.Teacher
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/teachers/{id} [put]
func UpdateTeacher(c *fiber.Ctx) error {
	db := database.DB
	var teacher model.Teacher
		if err := db.First(&teacher, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Teacher not found",
			})
		}

		if err := c.BodyParser(&teacher); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		if err := db.Save(&teacher).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update teacher",
			})
		}

		return c.JSON(teacher)
}

// DeleteStudentByID xóa một giáo viên dựa trên ID
// @Summary Xóa giáo viên
// @Description Xóa một giáo viên dựa trên ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int true "ID của giáo viên"
// @Success 200 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /users/teachers/{id} [delete]
func DeleteTeacherByID(c *fiber.Ctx) error {
	db := database.DB
	var teacher model.Teacher
		if err := db.First(&teacher, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Teacher not found",
			})
		}

		if err := db.Delete(&teacher).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete teacher",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Teacher deleted",
		})
}


// func Test(c *fiber.Ctx) error {
// 	// Bắt đầu một giao dịch
// 	db := database.DB
// 	tx := db.Begin()

// 	// Tạo một bản ghi mới
// 	teacher := model.Teacher{
// 		ID : 11,
// 		Name: "Thuyen",
// 		Subject: "Vat ly",
// 	}
// 	if err := tx.Create(&teacher).Error; err != nil {
// 		// Nếu có lỗi xảy ra, hoàn tác giao dịch và xử lý lỗi
// 		tx.Rollback()
// 		fmt.Println("Failed to create teacher:", err)
// 		return err
// 	}

// 	// // Đọc dữ liệu
// 	// var students []model.Teacher
// 	// if err := tx.Find(&students).Error; err != nil {
// 	// 	// Nếu có lỗi xảy ra, hoàn tác giao dịch và xử lý lỗi
// 	// 	tx.Rollback()
// 	// 	fmt.Println("Failed to retrieve students:", err)
// 	// 	return err
// 	// }

// 	// Cập nhật một bản ghi
// 	if err := tx.Model(&teacher).Update("Age", 21).Error; err != nil {
// 		// Nếu có lỗi xảy ra, hoàn tác giao dịch và xử lý lỗi
// 		tx.Rollback()
// 		fmt.Println("Failed to update teacher:", err)
// 		return err
// 	}

// 	// // Xóa một bản ghi
// 	// if err := tx.Delete(&teacher).Error; err != nil {
// 	// 	// Nếu có lỗi xảy ra, hoàn tác giao dịch và xử lý lỗi
// 	// 	tx.Rollback()
// 	// 	fmt.Println("Failed to delete teacher:", err)
// 	// 	return err
// 	// }

// 	// Commit giao dịch nếu không có lỗi
// 	if err := tx.Commit().Error; err != nil {
// 		// Nếu có lỗi xảy ra, hoàn tác giao dịch và xử lý lỗi
// 		tx.Rollback()
// 		fmt.Println("Failed to commit transaction:", err)
// 		return err
// 	}
// 	return tx.Commit().Error
// }