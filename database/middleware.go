package database

import (
	"net/http"

	"github.com/gorilla/context"
)

const key = "database"

// Middleware adds the database Service to this context if it supports
// the Setter interface.
func (d *Client) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, key, d)
		next.ServeHTTP(w, r)
	})
}
