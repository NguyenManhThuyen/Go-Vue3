package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	return c.JSON("OK")
}

func Wellcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"mess" : "Wellcome to CSV_SCHOOL"})
}
