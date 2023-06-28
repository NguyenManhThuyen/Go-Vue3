package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DataResponse struct {
	Status    bool          `json:"status" `
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
	ValidateError interface{} `json:"validate_error"`
}

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}
