package middleware

import (
	"api/config"
	"api/database"
	"api/modules/user/model"
	"api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Check Key App
func AppInfo(c *fiber.Ctx) error {
	app_key := c.Get("appkey")

	if len(app_key) == 0 || app_key != config.Config("APP_KEY") {
		return c.Status(400).JSON(fiber.Map{
			"message": "System is not support Key",
		})
	}

	return c.Next()
}

// Authen
func AppAuthen(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	store := database.ConfigSession()
	is_error := 0

	// Check token valid
	tokenData, err := utils.ExtractTokenData(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check user exists
	if res := db.Where("username = ? and deleted_at is null", tokenData.Username).First(&user); res.RowsAffected <= 0 {
		is_error = 1
	}

	// Check session Exist and comparse token
	sess, err := store.Get(tokenData.Username)
	authen := strings.Split(c.Get("token"), " ")
	if err != nil || len(sess) == 0 || string(sess) != authen[1] {
		is_error = 1
	}

	// Check User Agent / IP
	ip_address := c.IP()
	if ip_address != tokenData.IPAdress {
		is_error = 1
	}

	if is_error == 1 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Token is incorrect",
		})
	}

	return c.Next()
}

// SupberAdmin
func SuperAdmin(c *fiber.Ctx) error {
	// Check token valid
	tokenData, err := utils.ExtractTokenData(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if tokenData.Username == "csvadmin" || tokenData.Username == "phuxuan" {
		return c.Next()
	} else {
		return c.Status(500).JSON(fiber.Map{"message": "Permission deny"})
	}
}
