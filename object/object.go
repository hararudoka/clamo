package object

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// User will be used to store user information in DB and in JSON responses
type User struct {
	ID uuid.UUID `json:"id"` // unique id for user

	Username string `json:"username"` // unique name that user will be able to change

	Password string `json:"password"` // just password for authentication
}

func (u User) ToJSON() []byte {
	json, err := json.Marshal(u)
	if err != nil {
		return nil
	}
	return json
}

func (u *User) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &u)
	if err != nil {
		return err
	}
	return nil
}

type Message struct {
	ID uuid.UUID `json:"id"` // unique id for message

	FromID int `json:"from_id"` // id of user who sent the message
	ToID   int `json:"to_id"`   // id of user who will receive the message

	Text string `json:"text"` // content of the message
}

// List of errors that can be returned by something in this package
var (
	ErrNotFound         = errors.New("not found")
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrTakenUsername    = errors.New("username is already taken")

	ErrUsernameNotSpecified = errors.New("username is not specified")
	ErrPassNotSpecified     = errors.New("password is not specified")
	ErrDataNotSpecified     = errors.New("data is not specified")

	ErrWrongID      = errors.New("wrong id")
	ErrNotFoundUser = errors.New("not found user")
)
