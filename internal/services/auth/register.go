package auth

import (
	"errors"
	"ticket-backend/internal/database"
	"ticket-backend/internal/models/auth"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	db *gorm.DB
}

func NewRegisterService() *RegisterService {
	return &RegisterService{
		db: database.DB,
	}
}

func (s *RegisterService) Register(req *auth.RegisterRequest) error {
	var count int64
	s.db.Model(&auth.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &auth.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Name:         req.Name,
		Email:        req.Email,
		Role:         "it_support",
	}

	return s.db.Create(user).Error
}
