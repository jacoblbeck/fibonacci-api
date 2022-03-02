package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Health represents the API handler to
// report the health status for the server.
func Health(w http.ResponseWriter, r *http.Request) {
	logrus.Info("checking health")

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("ok"))
	if err != nil {
		logrus.Error(err)
	}
}
