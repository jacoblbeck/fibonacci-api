package database

import (
	"net/http"
)

// Service represents the interface for Marina integrating
// with the different supported Database backends.
type Service interface {
	// Middleware defines a function that
	// adds a registry client to the middleware
	Middleware(http.Handler) http.Handler

	GetCurrent() int

	SetCurrent(int) error

	GetPrevious() int

	SetPrevious(int) error
}
