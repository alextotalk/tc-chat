package service

import (
	"context"
	"log"
	"time"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/storage"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// TODO: create middleware file/folder for JWT
type AuthService struct {
	Service
	repo      *storage.UserRepository
	secretKey []byte
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
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
	tokenString, err := s.CreateToken(email)
	log.Print("Token created: ", tokenString)
	if err != nil {
		return nil, err
	}
	if hashedPassword != user.Password {
		return nil, log.Output(1, "Incorrect password. Please try again")
	}
	return user, nil
}

func (s *AuthService) CreateToken(email string) (string, error) {
	s.secretKey = make([]byte, 32)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return log.Output(1, "Error: token is not valid")
	}

	return nil
}
