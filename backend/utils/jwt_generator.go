package utils

import (
	"api/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(username, user_agent, ip_address string) (string, error) {
	secret := config.Config("JWT_SECRET_KEY")
	time_expire := config.Config("JWT_EXPIRED_TIME")

	minutesCount, _ := strconv.Atoi(time_expire)

	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["useragent"] = user_agent
	claims["ipaddress"] = ip_address
	claims["createdat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateAccessTokenMobile(username, user_agent, ip_address string, permission int) (string, error) {
	secret := config.Config("JWT_SECRET_KEY")
	time_expire := config.Config("JWT_EXPIRED_TIME")

	minutesCount, _ := strconv.Atoi(time_expire)

	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["useragent"] = user_agent
	claims["ipaddress"] = ip_address
	claims["permission"] = permission
	claims["createdat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func EncodeDataTokenMobile(employee_id string, date_in_seconds int64, coordinates string, shift_id int) (string, error) {
	secret := config.Config("JWT_DATA_SECRET_KEY")
	time_expire := config.Config("JWT_DATA_EXPIRED_TIME")

	minutesCount, _ := strconv.Atoi(time_expire)

	claims := jwt.MapClaims{}

	claims["employee_id"] = employee_id
	claims["date_in_seconds"] = date_in_seconds
	claims["coordinates"] = coordinates
	claims["shift_id"] = shift_id
	claims["createdat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
