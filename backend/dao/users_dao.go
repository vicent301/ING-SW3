package dao

import (
	"backend/database"
	"backend/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// 🔧 AutoMigrar tabla
func AutoMigrateUser() {
	database.DB.AutoMigrate(&User{})
}

// 🔐 Crear usuario con password hasheado
func CreateUser(u domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	entity := User{
		Name:     u.Name,
		Email:    u.Email,
		Password: string(hashedPassword),
	}
	return database.DB.Create(&entity).Error
}

// 🔍 Buscar usuario por email
func GetUserByEmail(email string) (*domain.User, error) {
	var e User
	if err := database.DB.Where("email = ?", email).First(&e).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		ID:       e.ID,
		Name:     e.Name,
		Email:    e.Email,
		Password: e.Password,
	}, nil
}

// 🔍 Buscar usuario por ID
func GetUserByID(id uint) (*domain.User, error) {
	var e User
	if err := database.DB.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		ID:    e.ID,
		Name:  e.Name,
		Email: e.Email,
	}, nil
}
