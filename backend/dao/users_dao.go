package dao

import (
	"backend/database"
	"backend/domain"
	"errors"

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
	if database.DB == nil {
		return // evita panic si DB no está inicializada en test
	}
	database.DB.AutoMigrate(&User{})
}

// 🔐 Crear usuario con password hasheado
func CreateUser(u domain.User) error {
	if database.DB == nil {
		return errors.New("database not initialized")
	}

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
	if database.DB == nil {
		return nil, errors.New("database not initialized")
	}

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
	if database.DB == nil {
		return nil, errors.New("database not initialized")
	}

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
