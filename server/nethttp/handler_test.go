package nethttp

import (
	"testing"

	"github.com/hararudoka/clamo/object"
	"github.com/hararudoka/clamo/server/service"
)

// Mock of the DB
type MockStore struct {
	user    map[int]object.User
	message map[int]object.Message
}

// User -> save
func (ms MockStore) SaveUser(object.User) error {
	return nil
}

// Message -> save
func (ms MockStore) SaveMessage(object.Message) error {
	return nil
}

// User.id -> User
func (ms MockStore) GetUser(int) (object.User, error) {
	return object.User{}, nil
}

// Message.id -> Message
func (ms MockStore) GetMessage(int) (object.Message, error) {
	return object.Message{}, nil
}

// Username+Password -> User
func (ms MockStore) CheckLogin(string, string) (object.User, error) {
	return object.User{}, nil
}

// Main application test
func TestHandler_ServeHTTP(t *testing.T) {
	dbMock := MockStore{}
	s := service.New(dbMock)
	h := New(*s)

	tests := []struct {
		name   string
		path   string
		method string
		status int
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test condition

			_ = h
		})
	}
}
