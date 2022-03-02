package database

import (
	"github.com/jacoblbeck/fibonacci-api/types"
	"github.com/sirupsen/logrus"
)

type FibonacciService service

//GetCurrent gets the current value of the fibonacci sequence
func (svc *FibonacciService) GetCurrent() (int64, error) {
	logrus.Info("getting current from db")
	// variable to store query results
	f := types.Fibonacci{}

	db := svc.client.Database.Table("fibonaccis").Where("id = 1").Select("current").Scan(&f)

	if db.Error != nil {
		return -1, db.Error
	}

	return f.Current, nil
}

//SetCurrent sets the current value of the fibonacci sequence
func (svc *FibonacciService) SetCurrent(newCurrent int64) error {

	db := svc.client.Database.Table("fibonaccis").Where("id = 1").Update("current", newCurrent)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

//GetPrevious gets the previous value of the fibonacci sequence
func (svc *FibonacciService) GetPrevious() (int64, error) {
	// variable to store query results
	f := new(types.Fibonacci)

	// send query to the database and store result in variable
	db := svc.client.Database.Table("fibonaccis").Where("id = 1").Select("previous").Scan(&f)

	if db.Error != nil {
		return -1, db.Error
	}

	return f.Previous, nil
}

//SetPrevious sets the previous value of the fibonacci sequence
func (svc *FibonacciService) SetPrevious(newPrev int64) error {

	db := svc.client.Database.Table("fibonaccis").Where("id = 1").Update("previous", newPrev)

	if db.Error != nil {
		return db.Error
	}

	return nil
}
