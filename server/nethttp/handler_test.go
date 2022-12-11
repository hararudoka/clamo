package nethttp

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/hararudoka/clamo/object"
	"github.com/hararudoka/clamo/server/service"
)

// Mock of the DB
type MockStore struct {
	user    map[string]object.User
	message map[uuid.UUID]object.Message
}

func newMock() *MockStore {
	return &MockStore{
		user:    make(map[string]object.User),
		message: make(map[uuid.UUID]object.Message),
	}
}

// User -> save
func (ms *MockStore) SaveUser(user object.User) error {
	if _, ok := ms.user[user.Username]; ok {
		return object.ErrTakenUsername
	}
	ms.user[user.Username] = user
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

// // Main application test
// func TestHandler_ServeHTTP(t *testing.T) {
// 	dbMock := MockStore{}
// 	s := service.New(dbMock)
// 	h := New(*s)

// 	tests := []struct {
// 		name   string
// 		path   string
// 		method string
// 		status int
// 	}{
// 		{},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// test condition

// 			_ = h
// 		})
// 	}
// }

func TestGetUser(t *testing.T) {
	dbMock := newMock()
	s := service.New(dbMock)
	h := New(s)

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
			path:   "/register",
			method: http.MethodPost,
			status: http.StatusConflict,
			err:    errors.New(object.ErrTakenUsername.Error()),
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
			err:    errors.New(object.ErrCredentialsNotSpecified.Error()),
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

			status := w.Code

			// TODO: fix error condition.
			var actualError string
			if status != http.StatusOK {
				actualError = string(w.Body.Bytes())
			}

			// w.Body to object.User
			var user object.User
			if actualError == "" {
				err = user.FromJSON(w.Body.Bytes())
				if err != nil {
					t.Fatal(string(w.Body.Bytes()), err)
				}

				u, _ := s.SaveUser(user)
				user.ID = u
			}

			// check status code
			if status != tt.status {
				t.Errorf("handler returned wrong status code:\n got '%v'\nwant '%v'",
					status, tt.status)
			}

			// check error
			if tt.err != nil {
				if tt.err.Error() != actualError {
					t.Errorf("handler returned wrong error:\n got '%s'\nwant '%s'",
						actualError, tt.err)
				}
			} else {
				// check response body, uuid.ID is random, so we can't check it, but we can check is it empty
				if (user.Password != tt.expected.Password || user.Username != tt.expected.Username || user.ID == uuid.UUID{}) {
					t.Errorf("handler returned unexpected body:\n got '%v'\nwant '%v'",
						user, tt.expected)
				}
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
