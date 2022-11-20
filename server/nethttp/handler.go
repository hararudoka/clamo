package nethttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hararudoka/clamo/object"
	"github.com/hararudoka/clamo/server/service"
)

// Handler is a struct for http handlers, which contains Service
type Handler struct {
	service.Service
}

// New creates new Handler
func New(s service.Service) http.Handler {
	return Handler{
		s,
	}
}

// ServeHTTP is a root method for http handlers. TOODO: implement
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	// big conditional statement for routing
	if r.Method == http.MethodGet {
		if r.URL.Path == "/getUser" {
			h.GetUser(w, r)
			return
		}
		if r.URL.Path == "/getMessage" {
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
		if r.URL.Path == "/sendMessage" {
			h.SendMessage(w, r)
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

// Error is a handler for errors, returns error in JSON
func (h Handler) Error(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("usual error:", err, r.Method, r.URL.Path)

	statusCode := 500

	// error handling. TODO: think about how to implement this (LocalError{Code, GlobalError}????)
	switch err {
	case object.ErrNotFound:
		statusCode = 404
	case object.ErrTakenUsername:
		statusCode = 409
	case object.ErrIDNotSpecified, object.ErrWrongID:
		statusCode = 400
	}

	var resp struct {
		Error string `json:"error"`
	}

	// write error message to response
	resp.Error = err.Error()
	w.WriteHeader(statusCode)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("CRITICAL ERROR:", err, r.Method, r.URL.Path)
	}
}

// GetUser is a handler for GET /getUser, requests user id, returns user's info in JSON
func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
}

// GetMessage is a handler for GET /getMessage, requests message id, returns message's info in JSON
func (h Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
}

// Register is a handler for POST /register, requests user's username and password, returns user's info in JSON
func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
}

// Login is a handler for POST /login, requests user's username and password, returns user's info in JSON
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
}

// SendMessage is a handler for POST /sendMessage, requests a token, returns a message in JSON
func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
}
