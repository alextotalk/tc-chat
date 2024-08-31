package service

import (
	"context"
	"fmt"
	"log"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Service
	repo *storage.UserRepository
}

func (s *AuthService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *AuthService) Singup(ctx context.Context, name string, password string, email string, phone string) error {
	hashedPassword, _ := s.hashPassword(password)

	user := domain.User{Name: name, Email: email, Password: hashedPassword, Phone: phone}
	if err := s.repo.AppendUser(ctx, user); err != nil {
		return err
	}
	log.Println("Successfully added new user with email", email)
	return nil
}

func (s *AuthService) LoginUser(ctx context.Context, email string, password string) (*domain.User, error) {
	hashedPassword, _ := s.hashPassword(password)
	user, err := s.repo.ByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if hashedPassword != user.Password {
		return nil, fmt.Errorf("Incorrect password. Please try again")
	}
	return user, nil
}
