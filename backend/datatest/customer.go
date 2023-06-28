package datatest

import (
	"api/database"
	"api/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateCustomer(c *fiber.Ctx) error {
	number := 30

	for i := 1; i <= number; i++ {
		x := i + 10
		db := database.DB
		customer := new(model.Customer)

		customer.Company = fmt.Sprint("CUBE SYSTEM VIETNAM ", x)
		customer.Representative = fmt.Sprint("Customer test ", x)
		customer.Phone = "028 3715 5701"
		customer.Address = "86/53 TCH 36, Quận 12, HCM"
		db.Create(&customer)

	}

	return c.JSON(fiber.Map{"message": "Datatest created"})
}
