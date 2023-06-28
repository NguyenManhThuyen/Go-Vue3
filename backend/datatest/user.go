package datatest

import (
	"api/database"
	"api/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Createuser(c *fiber.Ctx) error {
	number := 30

	for i := 1; i <= number; i++ {
		x := i + 10
		db := database.DB
		user := new(model.User)
		userprofile := new(model.UserProfile)

		username := fmt.Sprint("admin", x)
		if res := db.Where("username = ?", username).Find(&user); res.RowsAffected <= 0 {
			user.Username = username
			user.Password = model.HashPassword(username)
			db.Create(&user)

			userprofile.UserId = user.ID
			userprofile.Name = fmt.Sprint("User test ", x)
			userprofile.Birthday = "2008-03-18"
			userprofile.Phone = "028 3715 5701"
			userprofile.Avatar = "assets/images/profiles/test.png"
			db.Create(&userprofile)

		}
	}

	return c.JSON(fiber.Map{"message": "Datatest created"})
}
