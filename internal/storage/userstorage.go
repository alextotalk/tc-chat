package storage

import (
	"container/list"
	"context"
	"time"

	"github.com/alextotalk/tc-chat/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	Repository
}

func (r *UserRepository) AllUsers(ctx *context.Context) (*list.List, error) {
	query := "SELECT * FROM Users"
	rows, err := r.db.Query(context.Background(), query)
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

func (r *UserRepository) ByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT * FROM user WHERE email = @email`
	args := pgx.NamedArgs{
		"email": email,
	}
	row := r.db.QueryRow(ctx, query, args)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.RegisteredAt, &user.LastVisitAt, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AppendUser(ctx context.Context, user domain.User) error {
	query := `INSERT INTO user VALUES (@name, @email, @phone, @password, @registered_at, @role)`
	args := pgx.NamedArgs{
		"name":          user.Name,
		"email":         user.Email,
		"phone":         user.Phone,
		"password":      user.Password,
		"registered_at": time.Now(),
		"role":          user.Role,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, currentEmail string, user domain.User) error {
	query := `UPDATE user
	SET (@name, @email, @phone, @password) WHERE email = ` + currentEmail
	args := pgx.NamedArgs{
		"name":     user.Name,
		"email":    user.Email,
		"phone":    user.Phone,
		"password": user.Password,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateRole(ctx context.Context, role string, email string) error {
	query := `UPDATE user SET (@role) WHERE email = ` + email
	args := pgx.NamedArgs{
		"role": role,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, email string) error {
	query := `DELETE FROM user WHERE email = @email`
	args := pgx.NamedArgs{
		"email": email,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}
