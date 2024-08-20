package service

import (
	"github.com/alextotalk/tc-chat/internal/storage"
)

type Service struct {
	repo *storage.Repository
}

func NewService(repos *storage.Repository) *Service {
	return &Service{
		repos,
	}
}
