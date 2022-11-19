package service

// Storage is a interface (list of requested methods) for storage (DB) layer. TODO: implement this
type Storage interface{}

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
