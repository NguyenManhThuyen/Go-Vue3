package RsaAlgorithm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Encrypt(app *fiber.App) {
	// Tạo một cặp khóa RSA mới
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Không thể tạo cặp khóa RSA:", err)
		return
	}

	// Lưu khóa riêng tư vào file
	privateKeyFile, err := os.Create("Algorithm/RSA_Algorithm/File/private.pem")
	if err != nil {
		fmt.Println("Không thể tạo file private.pem:", err)
		return
	}
	defer privateKeyFile.Close()

	// Chuyển đổi khóa riêng tư sang định dạng PEM
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		fmt.Println("Không thể lưu khóa riêng tư:", err)
		return
	}

	fmt.Println("Đã lưu khóa riêng tư vào file RSA_Algorithm/file/private.pem")

	// Lưu khóa công khai vào file
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create("Algorithm/RSA_Algorithm/file/public.pem")
	if err != nil {
		fmt.Println("Không thể tạo file RSA_Algorithm/file/public.pem:", err)
		return
	}
	defer publicKeyFile.Close()

	// Chuyển đổi khóa công khai sang định dạng PEM
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}
	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		fmt.Println("Không thể lưu khóa công khai:", err)
		return
	}

	fmt.Println("Đã lưu khóa công khai vào file RSA_Algorithm/file/public.pem")

	// Đọc dữ liệu từ file nguồn
	data, err := ioutil.ReadFile("Algorithm/input.txt")
	if err != nil {
		fmt.Println("Không thể đọc dữ liệu từ input.txt:", err)
		return
	}

	// Mã hóa dữ liệu bằng khóa công khai
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		fmt.Println("Không thể mã hóa dữ liệu:", err)
		return
	}

	// Lưu dữ liệu đã mã hóa vào file
	err = ioutil.WriteFile("Algorithm/RSA_Algorithm/File/encrypted.bin", encryptedData, 0644)
	if err != nil {
		fmt.Println("Không thể lưu dữ liệu đã mã hóa:", err)
		return
	}

	fmt.Println("Đã lưu dữ liệu đã mã hóa vào file RSA_Algorithm/file/encrypted.bin")
}

func Decrypt(app *fiber.App) {
	// Đọc khóa riêng tư từ file
	privateKeyFile, err := os.Open("Algorithm/RSA_Algorithm/file/private.pem")
	if err != nil {
		fmt.Println("Không thể đọc khóa riêng tư từ file private.pem:", err)
		return
	}
	defer privateKeyFile.Close()

	privateKeyBytes, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		fmt.Println("Không thể đọc khóa riêng tư từ file:", err)
		return
	}

	privateKeyPEM, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPEM.Bytes)
	if err != nil {
		fmt.Println("Không thể phân tích cú pháp khóa riêng tư:", err)
		return
	}

	// Đọc dữ liệu đã mã hóa từ file
	encryptedData, err := ioutil.ReadFile("Algorithm/RSA_Algorithm/File/encrypted.bin")
	if err != nil {
		fmt.Println("Không thể đọc dữ liệu đã mã hóa từ file encrypted.bin:", err)
		return
	}

	// Giải mã dữ liệu bằng khóa riêng tư
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		fmt.Println("Không thể giải mã dữ liệu:", err)
		return
	}

	// Lưu dữ liệu đã giải mã vào file
	err = ioutil.WriteFile("Algorithm/decrypted.txt", decryptedData, 0644)
	if err != nil {
		fmt.Println("Không thể lưu dữ liệu đã giải mã:", err)
		return
	}

	fmt.Println("Đã lưu dữ liệu đã giải mã vào file decrypted.txt")
}
