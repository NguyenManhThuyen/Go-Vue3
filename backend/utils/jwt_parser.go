package utils

import (
	"api/config"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenData struct {
	Username  string
	Useragent string
	IPAdress  string
	Createdat int64
	Expires   int64
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("x-csv-token")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 && onlyToken[0] == "Bearer" {
		return onlyToken[1]
	}

	return ""
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	if len(tokenString) == 0 {
		msg := config.GetMessageCode("TOKEN_INCORRECT")
		return nil, errors.New(msg)
	}

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Config("JWT_SECRET_KEY")), nil
}

func ExtractTokenData(c *fiber.Ctx) (*TokenData, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &TokenData{
			Username:  fmt.Sprint(claims["username"]),
			Createdat: int64((claims["iat"]).(float64)),
			Expires:   int64((claims["exp"]).(float64)),
		}, nil
	}

	return nil, err
}

/** Mobile */

type TokenDataMobile struct {
	Username   string
	Useragent  string
	IPAdress   string
	Permission int
	Createdat  int64
	Expires    int64
}

func ExtractTokenDataMobile(c *fiber.Ctx) (*TokenDataMobile, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &TokenDataMobile{
			Username:   fmt.Sprint(claims["username"]),
			Useragent:  fmt.Sprint(claims["useragent"]),
			IPAdress:   fmt.Sprint(claims["ipaddress"]),
			Permission: int(claims["permission"].(float64)),
			Createdat:  int64((claims["createdat"]).(float64)),
			Expires:    int64((claims["exp"]).(float64)),
		}, nil
	}

	return nil, err
}

/**
EncodedData
*/

type Data struct {
	EmployeeId    string `json:"employee_id"`
	DateInSeconds int64  `json:"date_in_seconds"`
	Coordinates   string `json:"coordinates"`
	ShiftId       int64  `json:"shift_id"`
}

func DecodeData(encodedData string) (*Data, error) {
	if len(encodedData) == 0 {
		return nil, errors.New("data is incorrect")
	}

	data, err := jwt.Parse(encodedData, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("JWT_DATA_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := data.Claims.(jwt.MapClaims)
	if ok && data.Valid {
		return &Data{
			EmployeeId:    fmt.Sprint(claims["employee_id"]),
			DateInSeconds: int64((claims["date_in_seconds"]).(float64)),
			Coordinates:   fmt.Sprint(claims["coordinates"]),
			ShiftId:       int64((claims["shift_id"]).(float64)),
		}, nil
	}

	return nil, err
}
