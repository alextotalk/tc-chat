package service

import (
	"context"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/storage"
)

type ProfileService struct {
	Service
	repo *storage.UserRepository
}

func (s *ProfileService) EditProfile(ctx context.Context, currentEmail string, name string, newEmail string,
	phone string, password string) error {

	user := domain.User{Name: name, Email: newEmail, Phone: phone, Password: password}
	err := s.repo.UpdateUser(ctx, currentEmail, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, email string) error {
	err := s.repo.DeleteUser(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
