package nethttp

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/hararudoka/clamo/object"
	"github.com/hararudoka/clamo/server/service"
)

// Mock of the DB
type MockStore struct {
	user    map[uuid.UUID]object.User
	message map[uuid.UUID]object.Message
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

func TestGetUser(t *testing.T) {
	dbMock := MockStore{}
	s := service.New(dbMock)
	h := New(*s)

	tests := []struct {
		name     string
		input    object.User
		expected object.User
		path     string
		method   string
		status   int
		err      error
	}{
		{
			name: "Create User",
			input: object.User{
				Username: "aboba",
				Password: "1",
			},
			expected: object.User{
				ID:       mustUUID(),
				Username: "aboba",
				Password: "1",
			},
			path:   "/register",
			method: http.MethodPost,
			status: http.StatusOK,
		},
		{
			name: "Create Same User",
			input: object.User{
				Username: "aboba",
				Password: "1",
			},
			expected: object.User{
				ID:       mustUUID(),
				Username: "aboba",
				Password: "1",
			},
			path:   "/register",
			method: http.MethodPost,
			status: http.StatusOK,
		},
		{
			name: "Create Empty User",
			input: object.User{
				Username: "",
				Password: "",
			},
			status: http.StatusBadRequest,
			path:   "/register",
			method: http.MethodPost,
			err:    object.ErrDataNotSpecified,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.input.ToJSON())

			// reader -> r
			r, err := http.NewRequest(tt.method, tt.path, reader)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.ServeHTTP)

			// r+w -> request
			handler.ServeHTTP(w, r)

			// check status code
			if status := w.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got '%v' want '%v'",
					status, tt.status)
			}

			// w.Body to object.User
			var user object.User
			err = user.FromJSON(w.Body.Bytes())
			if err != nil {
				t.Fatal(string(w.Body.Bytes()), err)
			}

			u, err := s.SaveUser(user)
			if err != nil {
				t.Fatal(err)
			}
			user.ID = u

			log.Println(tt.err)

			// check response body, uuid.ID is random, so we can't check it, but we can check is it empty
			if (user.Password != tt.expected.Password || user.Username != tt.expected.Username || user.ID == uuid.UUID{}) && tt.err == nil {
				t.Errorf("handler returned unexpected body: got '%v' want '%v'",
					user, tt.expected)
			}
		})
	}
}

func mustUUID() uuid.UUID {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return u
}
