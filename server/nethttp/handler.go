package nethttp

import (
	"net/http"

	"github.com/hararudoka/clamo/server/service"
)

// Service is a interface (list of requested methods) for service layer. TODO: implement this
type Service interface{}

// Handler is a struct for http handlers, which contains Service
type Handler struct {
	Service
}

// New creates new Handler
func New(s service.Service) http.Handler {
	return Handler{
		&s,
	}
}

// ServeHTTP is a root method for http handlers
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
