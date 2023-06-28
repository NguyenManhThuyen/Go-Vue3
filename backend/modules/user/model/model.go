package model

import (
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int         `json:"id" gorm:"primaryKey"`
	Username  string      `gorm:"index" json:"username" validate:"required"`
	Password  string      `json:"-" validate:"required,min=5" `
	Profile   UserProfile `json:"profile" gorm:"foreignKey:UserId"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"-"`
}

type UserProfile struct {
	ID              int        `json:"-" gorm:"primaryKey, index"`
	Name            string     `json:"name"`
	Birthday        string     `json:"birthday"`
	Phone           string     `json:"phone"`
	Avatar          string     `json:"avatar"`
	UserId          int        `json:"-"`
	Gender          string     `json:"gender"`
	Email           string     `json:"email"`
	Address         string     `json:"address"`
	DateJoin        string     `json:"date_join"`
	IdCard          string     `json:"id_card"`
	InsuranceNumber string     `json:"insurance_number"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"-"`
}

func (UserProfile) TableName() string {
	return "user_profile"
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
	cost := 9 // Min: 4, Max: 31
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes)
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (User) TableName() string {
	return "tbl_user"
}
