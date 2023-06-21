package main

import (
	// "api/RSA_Algorithm"
	// "api/AES_Algorithm"
	"api/RSA_Algorithm_File"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// RsaAlgorithm.Encrypt(app)
	// RsaAlgorithm.Decrypt(app)
	// AesAlgorithm.AesAlgorithm(app)
	RsaAlgorithmFile.Encrypt(app)
}
