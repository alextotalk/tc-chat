package service

import "github.com/alextotalk/atanika/internal/storage"

type Service struct {
	repo *storage.Repository
}

func NewService(repos *storage.Repository) *Service {
	return &Service{
		repos,
	}
}
