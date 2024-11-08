package storage

import (
	"container/list"
	"context"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserRoomRepository struct {
	Repository
}

func (r UserRepository) UsersByRoom(ctx context.Context, roomId int64) (*list.List, error) {
	query := `SELECT user_room.user_id FROM user_room
	RIGHT JOIN user ON user.id=user_room.user_id 
	WHERE user_room.room_id=@room_id`

	args := pgx.NamedArgs{
		"room_id": roomId,
	}
	rows, err := r.db.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := list.New()
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.RegisteredAt, &user.LastVisitAt, &user.Role)
		if err != nil {
			return nil, err
		}
		users.PushBack(user)
	}

	return users, nil
}

func (r UserRepository) Append(ctx context.Context, uRoom domain.UserRoom) error {
	query := `INSERT into user_room VALUES (@user_id, @room_id, @is_banned)`
	args := pgx.NamedArgs{
		"user_id":   uRoom.UserId,
		"room_id":   uRoom.RoomId,
		"is_banned": false,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) AllPairs(ctx context.Context) (*list.List, error) {
	query := `SELECT * FROM user_room`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pairs := list.New()
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.RegisteredAt, &user.LastVisitAt, &user.Role)
		if err != nil {
			return nil, err
		}
		pairs.PushBack(user)
	}

	return pairs, nil
}
