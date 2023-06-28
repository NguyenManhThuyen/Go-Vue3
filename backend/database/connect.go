package database

import (
	"api/config"
	"api/core"
	"fmt"
	"strconv"
	"time"
	"github.com/gofiber/storage/postgres"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Store *postgres.Storage

func Connect() bool {
	var err error
	status := true
	db_host := config.Config("DB_HOST")
	db_port := config.Config("DB_PORT")
	db_user := config.Config("DB_USER")
	db_password := config.Config("DB_PASSWORD")
	db_name := config.Config("DB_NAME")
	db_ssh := config.Config("DB_SSH")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db_host, db_port, db_user, db_password, db_name, db_ssh)

	DB, err = gorm.Open(postgresDriver.Open(dsn))
	if err != nil {
		status = false
		core.WriteLog("ERROR | DATABASE CONNECT")
	}

	if !status {
		return false
	}

	ConfigSession()

	return status
}

func ConfigSession() *postgres.Storage {
	host := config.Config("DB_HOST")
	port := config.Config("DB_PORT")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	name := config.Config("DB_NAME")
	sshmode := config.Config("SSH")
	post, _ := strconv.Atoi(port)

	store := postgres.New(postgres.Config{
		Host:       host,
		Port:       post,
		Username:   user,
		Password:   password,
		Database:   name,
		Table:      "session",
		Reset:      false,
		GCInterval: 10 * time.Second,
		SslMode:    sshmode,
	})

	return store

}