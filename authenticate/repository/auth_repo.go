package repository

import (
	"net/http"
	"users/apperror"
	"users/authenticate"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (authRepo *AuthRepository) GetDataLogin(username string) (*authenticate.Users, error) {
	err := authRepo.db.First(&authenticate.Users{}, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	var result *authenticate.Users
	authRepo.db.Where("username = ? ", username).Find(&authenticate.Users{}).Scan(&result)

	return result, nil

}

func (authRepo *AuthRepository) CreateUser(data *authenticate.Users) error {
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.New(http.StatusInternalServerError, "password tidak dapat di enkripsi")
	}

	data.Password = string(hashedPin)

	err = authRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil

}
