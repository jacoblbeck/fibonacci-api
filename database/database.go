package database

import (
	"net/http"
)

// Service represents the interface for fibonacci integrating
// with the different supported Database backends.
type Service interface {
	// Middleware defines a function that
	// adds a database client to the middleware
	Middleware(http.Handler) http.Handler

	GetCurrent() int

	SetCurrent(int) error

	GetPrevious() int

	SetPrevious(int) error
}
