package service

import (
	"context"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/alextotalk/tc-chat/internal/storage"
)

type MessageService struct {
	Service
	repo *storage.MessageRepository
}

func (s *MessageService) SendMessage(ctx context.Context, text string, senderId int64, roomId int64) error {
	message := domain.Message{SenderId: senderId, RoomId: roomId, Text: text}
	if err := s.repo.CreateMessage(ctx, message); err != nil {
		return err
	}
	return nil
}

func (s *MessageService) EditMessage(ctx context.Context, id int, text string) error {
	return s.repo.UpdateText(ctx, id, text)

}

func (s *MessageService) DeleteMessage(ctx context.Context, id int) error {
	return s.repo.DeleteMessage(ctx, id)
}
