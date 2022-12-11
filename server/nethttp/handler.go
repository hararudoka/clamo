package nethttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hararudoka/clamo/object"
	"github.com/hararudoka/clamo/server/service"
)

// Handler is a struct for http handlers, which contains Service
type Handler struct {
	*service.Service
}

// New creates new Handler
func New(s *service.Service) http.Handler {
	return Handler{
		s,
	}
}

// ServeHTTP is a root method for http handlers. TOODO: implement
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	// big conditional statement for routing
	if r.Method == http.MethodGet {
		if r.URL.Path == "/user" {
			h.GetUser(w, r)
			return
		}
		if r.URL.Path == "/message" {
			h.GetMessage(w, r)
			return
		}
		err = object.ErrNotFound
	} else if r.Method == http.MethodPost {
		if r.URL.Path == "/register" { // TODO: do not check auth here, but do it in all other routes
			h.Register(w, r)
			return
		}
		if r.URL.Path == "/login" { // TODO: do not check auth here, but do it in all other routes
			h.Login(w, r)
			return
		}
		if r.URL.Path == "/message" {
			h.PostMessage(w, r)
			return
		}
		err = object.ErrNotFound
	} else {
		err = object.ErrMethodNotAllowed
	}

	if err != nil {
		h.Error(w, r, err)
	}
}

func statusFromError(err error) int {
	switch err {
	case nil:
		return http.StatusOK
	case object.ErrNotFound:
		return http.StatusNotFound
	case object.ErrTakenUsername:
		return http.StatusConflict
	case object.ErrCredentialsNotSpecified, object.ErrWrongID:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

// Error is a handler for errors, returns error in plain text
func (h Handler) Error(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("usual error:", err, r.Method, r.URL.Path)

	var resp string

	// write error message to response
	resp = err.Error()
	w.WriteHeader(statusFromError(err))

	_, err = fmt.Fprint(w, resp)
	if err != nil {
		log.Println("CRITICAL ERROR:", err, r.Method, r.URL.Path)
	}
}

// GetUser is a handler for GET /user, requests user id, returns user's info in JSON
func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	h.Error(w, r, object.ErrNotImplemented)
}

// GetMessage is a handler for GET /message, requests message id, returns message's info in JSON
func (h Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	h.Error(w, r, object.ErrNotImplemented)
}

// Register is a handler for POST /register, requests user's username and password, returns user's info in JSON
func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	// get User from request
	var (
		err  error
		user object.User
	)
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	if user.Username == "" || user.Password == "" {
		h.Error(w, r, object.ErrCredentialsNotSpecified)
		return
	}

	id, err := h.Service.SaveUser(user)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	user.ID = id
	// write response
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		h.Error(w, r, err)
		return
	}
}

// Login is a handler for POST /login, requests user's username and password, returns user's info in JSON
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.Error(w, r, object.ErrNotImplemented)
}

// PostMessage is a handler for POST /message, requests a token, returns a message in JSON
func (h Handler) PostMessage(w http.ResponseWriter, r *http.Request) {
	h.Error(w, r, object.ErrNotImplemented)
}

// // base http requester
// func (s Ha) request(method, route string, body []byte) ([]byte, error) {
// 	req, err := http.NewRequest(method, DefaultApiURL+route, bytes.NewBuffer(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+s.token)
// 	req.Header.Set("Content-Type", "application/json")

// 	fmt.Println(req.Body)
// 	resp, err := s.client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(resp.Body)

// 	return buf.Bytes(), nil
// }
