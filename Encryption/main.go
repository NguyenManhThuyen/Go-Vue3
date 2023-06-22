package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func encryptData(c *fiber.Ctx) error {
	dataType := c.FormValue("dataType")
	algorithm := c.FormValue("encryptionAlgorithm")

	if dataType == "text" {
		inputText := c.FormValue("inputText")

		// Logic for encrypting text data using the selected algorithm
		// ...

		// Example response
		response := fmt.Sprintf("Text data encrypted using %s algorithm", algorithm)
		return c.SendString(response)
	} else if dataType == "file" || dataType == "folder" {
		// Retrieve uploaded file(s)
		form, err := c.MultipartForm()
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Error processing file(s)")
		}

		files := form.File["files"]

		// Create a temporary folder to store encrypted files
		tempFolder := "temp"
		if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
			os.Mkdir(tempFolder, os.ModePerm)
		}

		// Loop through each uploaded file
		for _, file := range files {
			// Open the uploaded file
			src, err := file.Open()
			if err != nil {
				log.Println(err)
				continue
			}
			defer src.Close()

			// Create the destination file path in the temporary folder
			fileExt := filepath.Ext(file.Filename)
			fileName := strings.TrimSuffix(file.Filename, fileExt)
			encryptedFileName := fileName + "_encrypted" + fileExt
			dstPath := filepath.Join(tempFolder, encryptedFileName)

			// Create the destination file
			dst, err := os.Create(dstPath)
			if err != nil {
				log.Println(err)
				continue
			}
			defer dst.Close()

			// Logic for encrypting the file using the selected algorithm
			// ...
			// Here, you can use your chosen encryption algorithm to encrypt the file
			// and write the encrypted content to the destination file

			// Example: Copy the content of the source file to the destination file
			_, err = io.Copy(dst, src)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		// Example response
		response := fmt.Sprintf("%d file(s) encrypted using %s algorithm", len(files), algorithm)
		return c.SendString(response)
	}

	return c.Status(http.StatusBadRequest).SendString("Invalid data type")
}

func decryptData(c *fiber.Ctx) error {
	encryptedData := c.FormValue("encryptedData")
	algorithm := c.FormValue("encryptionAlgorithm")

	// Logic for decrypting data using the selected algorithm
	// ...

	// Example response
	response := fmt.Sprintf("Data decrypted using %s algorithm", algorithm)
	return c.SendString(response)
}

func main() {
	app := fiber.New()

	// Endpoint for data encryption
	app.Post("/encrypt", encryptData)

	// Endpoint for data decryption
	app.Post("/decrypt", decryptData)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
