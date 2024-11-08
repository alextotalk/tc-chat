package storage

import (
	"context"
	"time"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/jackc/pgx/v5"
)

type MessageRepository struct {
	Repository
}

func (r *MessageRepository) MessageById(ctx context.Context, id int) (*domain.Message, error) {
	query := "SELECT * FROM message WHERE message.id= @id"
	args := pgx.NamedArgs{
		"id": id,
	}
	row := r.db.QueryRow(ctx, query, args)

	var message domain.Message
	err := row.Scan(&message.Id, &message.SenderId, &message.RoomId, &message.Text, &message.CreatedAt, &message.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessageRepository) CreateMessage(ctx context.Context, message domain.Message) error {
	query := "INSERT INTO message VALUES (@sender_id, @room_id, @text, @created_at, @updated_at)"
	//TODO:	make a method for finding suitable pair in  room-user

	args := pgx.NamedArgs{
		"sender_id":  message.SenderId,
		"room_id":    message.RoomId,
		"text":       message.Text,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) UpdateText(ctx context.Context, id int, text string) error {
	query := `UPDATE message SET (@text, @updated_at) WHERE id = @id`
	args := pgx.NamedArgs{
		"id":         id,
		"text":       text,
		"updated_at": time.Now(),
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessage(ctx context.Context, id int) error {
	query := `DELETE FROM message WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}
