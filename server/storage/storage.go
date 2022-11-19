package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// DB contains pgx connection to database
type DB struct {
	*pgx.Conn
}

// Open opens connection to database. It uses environment variables, so you should set them before calling this function.
func Open() (*DB, error) {
	name := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	mode := os.Getenv("DB_MODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", name, password, host, port, dbName, mode)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	log.Println("Connection!", connStr)

	// defer conn.Close(context.Background())
	return &DB{conn}, nil
}
