package controller

import (
	"api/config"
	"api/database"
	"api/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LoginParam struct {
	Username string `json:"username" `
	Password string `json:"password"`
}

// Login ================================================================================
// @Tags User
// @Summary Login
// @Description User login
// @Param body body LoginParam true " "
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /login [post]
// @Security ApiKeyAuth
func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	response := new(config.DataResponse)
	response.Status = false

	input := new(LoginInput)
	err := c.BodyParser(input)
	if err != nil || input == (&LoginInput{}) {
		response.Message = config.GetMessageCode("PARAM_ERROR")
		return c.JSON(response)
	}

	// check validate
	listCheck := []string{"username", "password"}
	item := map[string]string{"username": input.Username, "password": input.Password}
	errors := utils.RequireCheck(listCheck, item, map[string]string{})

	// Check max lenth
	listCheck = []string{"username:10", "password:20"}
	errors = utils.MaxLengthCheck(listCheck, item, errors)

	if len(errors) > 0 {
		response.Message = "validate"
		response.ValidateError = errors
		return c.JSON(response)
	}

	postBody, _ := json.Marshal(map[string]string{
		"username": input.Username,
		"password": input.Password,
	})

	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", config.Config("SSO_BASE_URL")+"/api/auth/login", responseBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-csv-app-id", config.Config("SSO_APP_ID"))
	req.Header.Add("x-csv-app-type", config.Config("SSO_APP_TYPE"))

	client := &http.Client{}
	resp, error := client.Do(req)
	if error != nil {
		fmt.Println(error)
	}
	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	jsonString := []byte(body)
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(jsonString), &data)

	resultData, _ := data["data"].(map[string]interface{})

	message := data["message"]
	statusCode := resp.StatusCode
	status := false

	if statusCode == 200 {
		status = true
		// Save session
		store := database.Store
		time_expire := config.Config("JWT_EXPIRED_TIME")
		minutesCount, _ := strconv.Atoi(time_expire)
		store.Set(input.Username, []byte(resultData["token"].(string)), time.Duration(minutesCount)*time.Minute)
	}

	// Return response
	response.Status = status
	response.Message = fmt.Sprintf("%v", message)
	response.Data = resultData
	return c.JSON(response)
}

// Check token ================================================================================
func CheckToken(c *fiber.Ctx) error {
	response := new(config.DataResponse)
	response.Status = true

	// Return response
	return c.JSON(response)
}
