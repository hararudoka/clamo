package service

import "github.com/hararudoka/clamo/object"

// Storage is a interface (list of requested methods) for storage (DB) layer. TODO: implement this
type Storage interface {
	// User -> save
	SaveUser(object.User) error
	// Message -> save
	SaveMessage(object.Message) error
	// User.id -> User
	GetUser(int) (object.User, error)
	// Message.id -> Message
	GetMessage(int) (object.Message, error)
	// Username+Password -> User
	CheckLogin(string, string) (object.User, error)
}

// Service is a struct for service layer
type Service struct {
	db Storage
}

// New creates new Service
func New(db Storage) *Service {
	return &Service{
		db: db,
	}
}
