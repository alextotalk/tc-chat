package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(pg *pgxpool.Pool) *Repository {
	return &Repository{
		pg,
	}
}
