package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/jacoblbeck/fibonacci-api/database"
	"github.com/sirupsen/logrus"
)

//GetNext is the api function to get the next value in the fibonacci sequence.
func GetNext(w http.ResponseWriter, r *http.Request) {
	logrus.Info("getting next")

	database := context.Get(r, "database").(*database.Client)

	cur, err := database.Fibonacci.GetCurrent()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		logrus.Error(err)
		return
	}

	if cur == 0 {
		database.Fibonacci.SetCurrent(1)
		database.Fibonacci.SetPrevious(0)

		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("1"))

		if err != nil {
			logrus.Error(err)
		}
		return
	} else {
		prev, err := database.Fibonacci.GetPrevious()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		next := cur + prev

		database.Fibonacci.SetPrevious(cur)
		database.Fibonacci.SetCurrent(next)

		w.WriteHeader(http.StatusOK)

		_, err = w.Write([]byte(strconv.FormatInt(next, 10)))
		if err != nil {
			logrus.Error(err)
		}
		return
	}

}

//GetCurrent is the api function to get the current value of the fibonacci sequence.
func GetCurrent(w http.ResponseWriter, r *http.Request) {
	logrus.Info("getting current")

	database := context.Get(r, "database").(*database.Client)

	cur, err := database.Fibonacci.GetCurrent()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(strconv.FormatInt(cur, 10)))
	if err != nil {
		logrus.Error(err)
	}
}

//GetPrevious is the api function to get the previous value of the fibonacci sequence.
func GetPrevious(w http.ResponseWriter, r *http.Request) {
	logrus.Info("getting previous")

	database := context.Get(r, "database").(*database.Client)

	prev, err := database.Fibonacci.GetPrevious()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(strconv.FormatInt(prev, 10)))
	if err != nil {
		logrus.Error(err)
	}
}

//ResetSequence sets the sequence back to 0.
func ResetSequence(w http.ResponseWriter, r *http.Request) {
	logrus.Info("resetting sequence")

	database := context.Get(r, "database").(*database.Client)

	database.Fibonacci.SetCurrent(0)
	database.Fibonacci.SetPrevious(0)

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("reset"))
	if err != nil {
		logrus.Error(err)
	}

}
