package model

import (
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int        `json:"id" gorm:"primaryKey, index"`
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required,min=5" `
	Language   string     `json:"language"`
	Permission int        `json:"permission"`
	LogVersion int        `json:"log_version"`
	CreatedId  int        `json:"created_id" gorm:"default:null"`
	CreatedAt  *time.Time `json:"created_at" gorm:"default:null"`
	UpdatedId  int        `json:"updated_id" gorm:"default:null"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedId  int        `json:"deleted_id" gorm:"default:null"`
	DeletedAt  *time.Time `json:"-"`
}

func (User) TableName() string {
	return "tbl_user"
}

type ErrorResponseuser struct {
	Field string
	Tag   string
	Value string
}

func ValidateUser(user User) []*ErrorResponseuser {
	var errors []*ErrorResponseuser
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponseuser
			element.Field = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func HashPassword(password string) string {
	cost := 7 // Min: 4, Max: 31
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes)
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
