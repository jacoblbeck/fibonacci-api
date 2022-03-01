package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Next(w http.ResponseWriter, r *http.Request) {
	logrus.Info("getting next")

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("next"))
	if err != nil {
		logrus.Error(err)
	}
}

func Current(w http.ResponseWriter, r *http.Request) {
	logrus.Info("getting current")

	w.WriteHeader(http.StatusNotImplemented)

	_, err := w.Write([]byte("Not Implemented"))
	if err != nil {
		logrus.Error(err)
	}
}

func Previous(w http.ResponseWriter, r *http.Request) {
	logrus.Info("getting previous")

	w.WriteHeader(http.StatusNotImplemented)

	_, err := w.Write([]byte("Not Implemented"))
	if err != nil {
		logrus.Error(err)
	}
}
