package pg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

func New(cfg Config) (*pgxpool.Pool, error) {

	dbURL := "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode

	// Pool configuration
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}
	//Tune pool settings
	config.MaxConnLifetime = 0 // Connections are not closed based on age
	config.MaxConnIdleTime = 0 // Connections are not closed based on idle time
	config.MaxConns = 10       // Adjust based on your workload

	// Create a new pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Check connection
	//if err := pool.Ping(context.Background()); err != nil {
	//	log.Fatalf("Unable to ping database: %v\n", err)
	//}

	return pool, nil
}
