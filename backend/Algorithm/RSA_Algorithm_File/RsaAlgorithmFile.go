package RsaAlgorithmFile

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
	privateKeyFile, err := os.Create("Algorithm/RSA_Algorithm_File/FileTestInput/private.pem")
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

	fmt.Println("Đã lưu khóa riêng tư vào file private.pem")

	// Lưu khóa công khai vào file
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create("Algorithm/RSA_Algorithm_File/FileTestInput/public.pem")
	if err != nil {
		fmt.Println("Không thể tạo file public.pem:", err)
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

	fmt.Println("Đã lưu khóa công khai vào file public.pem")

	// Đường dẫn đến thư mục cần mã hóa
	inputDir := "Algorithm"

	// Đường dẫn đến thư mục đầu ra cho các file đã mã hóa
	outputDir := "Algorithm/RSA_Algorithm_File/FileTestOutput"
	err = os.Mkdir(outputDir, 0755)
	if err != nil {
		fmt.Println("Không thể tạo thư mục đầu ra:", err)
		return
	}

	// Lấy danh sách các file trong thư mục đầu vào
	fileList, err := filepath.Glob(filepath.Join(inputDir, "*"))
	if err != nil {
		fmt.Println("Không thể đọc danh sách các file:", err)
		return
	}

	// Mã hóa từng file
	for _, file := range fileList {
		// Đọc dữ liệu từ file nguồn
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Không thể đọc dữ liệu từ file:", err)
			continue
		}

		// Mã hóa dữ liệu bằng khóa công khai
		encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
		if err != nil {
			fmt.Println("Không thể mã hóa dữ liệu:", err)
			continue
		}

		// Tạo đường dẫn đầu ra cho file đã mã hóa
		encryptedFilePath := filepath.Join(outputDir, filepath.Base(file))

		// Lưu dữ liệu đã mã hóa vào file
		err = ioutil.WriteFile(encryptedFilePath, encryptedData, 0644)
		if err != nil {
			fmt.Println("Không thể lưu dữ liệu đã mã hóa:", err)
			continue
		}

		fmt.Printf("Đã mã hóa file %s thành công\n", file)
	}

	fmt.Println("Đã mã hóa các file trong thư mục thành công")

	// Giải mã từng file
	for _, file := range fileList {
		// Đọc dữ liệu đã mã hóa từ file
		encryptedData, err := ioutil.ReadFile(filepath.Join(outputDir, filepath.Base(file)))
		if err != nil {
			fmt.Println("Không thể đọc dữ liệu đã mã hóa:", err)
			continue
		}

		// Giải mã dữ liệu bằng khóa riêng tư
		decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
		if err != nil {
			fmt.Println("Không thể giải mã dữ liệu:", err)
			continue
		}

		// Tạo đường dẫn đầu ra cho file đã giải mã
		decryptedFilePath := filepath.Join(outputDir, "decrypted_"+filepath.Base(file))

		// Lưu dữ liệu đã giải mã vào file
		err = ioutil.WriteFile(decryptedFilePath, decryptedData, 0644)
		if err != nil {
			fmt.Println("Không thể lưu dữ liệu đã giải mã:", err)
			continue
		}

		fmt.Printf("Đã giải mã file %s thành công\n", file)
	}

	fmt.Println("Đã giải mã các file trong thư mục thành công")
}
