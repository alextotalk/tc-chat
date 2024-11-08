package storage

import (
	"container/list"
	"context"
	"time"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/jackc/pgx/v5"
)

type RoomRepository struct {
	Repository
}

func (r *RoomRepository) AllRooms(ctx context.Context) (*list.List, error) {
	query := `SELECT * FROM room`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := list.New()
	for rows.Next() {
		var room domain.Room
		err := rows.Scan(&room.Id, &room.Title, &room.Description, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms.PushBack(room)
	}

	return rooms, nil

}

func (r *RoomRepository) FindById(ctx context.Context, id int) (*domain.Room, error) {
	query := `SELECT * FROM room WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	row := r.db.QueryRow(ctx, query, args)
	var room domain.Room
	err := row.Scan(&room.Id, &room.Title, &room.Description, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room domain.Room) error {
	query := `INSERT INTO room VALUES (@title, @description, @created_at, @updated_at)`
	args := pgx.NamedArgs{
		"title":       room.Title,
		"description": room.Description,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) UpdateRoom(ctx context.Context, room domain.Room) error {
	query := `UPDATE room SET (@title, @description, @updated_at)`
	args := pgx.NamedArgs{
		"title":       room.Title,
		"description": room.Description,
		"updted_at":   time.Now(),
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}
