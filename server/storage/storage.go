package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hararudoka/clamo/object"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

// User -> save
func (db *DB) SaveUser(user object.User) error {
	_, err := db.Exec(context.Background(), "INSERT INTO \"user\" (id, username, password) VALUES ($1, $2, $3)", user.ID, user.Username, user.Password)
	if isUniqueViolationErr(err) {
		return object.ErrTakenUsername
	}
	return err
}

// Message -> save
func (db *DB) SaveMessage(message object.Message) error {
	_, err := db.GetUser(message.ToID)
	if err != nil {
		return object.ErrNotFoundUser
	}
	row := db.QueryRow(context.Background(), "INSERT INTO \"message\" (from_id, to_id, text) VALUES ($1, $2, $3) RETURNING id", message.FromID, message.ToID, message.Text)
	err = row.Scan(&message.ID)
	return err
}

// User.id -> User
func (db *DB) GetUser(id int) (object.User, error) {
	var user object.User
	err := db.QueryRow(context.Background(), "SELECT * FROM \"user\" WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

// Message.id -> Message
func (db *DB) GetMessage(id int) (object.Message, error) {
	var message object.Message

	err := db.QueryRow(context.Background(), "SELECT * FROM \"message\" WHERE id = $1", id).Scan(&message.ID, &message.FromID, &message.Text)
	return message, err
}

// Username+Password -> User
func (db *DB) CheckLogin(username, password string) (object.User, error) {
	var user object.User

	err := db.QueryRow(context.Background(), "SELECT * FROM \"user\" WHERE username = $1 AND password = $2", username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return object.User{}, err
	}

	return user, nil
}

func isUniqueViolationErr(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return true
		}
	}
	return false
}
