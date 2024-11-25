package service

import (
	"context"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/storage"
)

type RoomService struct {
	Service
	repo *storage.RoomRepository
}

func (s RoomService) CreateRoom(ctx context.Context, title string, description string) error {
	room := domain.Room{Title: title, Description: description}
	if err := s.repo.CreateRoom(ctx, room); err != nil {
		return err
	}
	return nil
}
