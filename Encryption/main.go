package main

import (
	"api/AES_Algorithm"
	"api/RSA_Algorithm"
	"api/RSA_Algorithm_File"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("----------------------------------/n")
	RsaAlgorithm.Encrypt(app)
	fmt.Println("----------------------------------/n")
	RsaAlgorithm.Decrypt(app)
	fmt.Println("----------------------------------/n")
	AesAlgorithm.AesAlgorithm(app)
	fmt.Println("----------------------------------/n")
	RsaAlgorithmFile.Encrypt(app)
}
