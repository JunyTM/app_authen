package service

import (
	"errors"
	"ms_auth/infrastructure"
	"ms_auth/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Login(email string, password string) (*model.User, error)
	CreateUser(name string, email string, password string) (*model.User, error)
}

type userService struct {
	db *gorm.DB
}

func (s *userService) Login(email string, password string) (*model.User, error) {
	var user *model.User
	if err := s.db.Model(&model.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("=> User not found in database")
	}

	if err := ComparePassword(password, user.Password); err != nil {
		return nil, errors.New("=> Password not match")
	}

	return user, nil
}

func (s *userService) CreateUser(name string, email string, password string) (*model.User ,error) {
	hashedPassword, err := HashAndSolatPassword(password)
	if err != nil {
		return nil, err
	}
	newRecord := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	if err := s.db.Create(&newRecord).Error; err!= nil {
        return nil, err
    }
	return newRecord, nil
}

func NewUserService() UserService {
	db := infrastructure.GetDB()
	return &userService{
		db: db,
	}
}

func ComparePassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("Password not match")
	}
	return nil
}

func HashAndSolatPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}