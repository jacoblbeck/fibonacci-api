package database

import (
	"github.com/jacoblbeck/fibonacci-api/types"
	"github.com/sirupsen/logrus"
)

type FibonacciService service

//GetCurrent gets the current value of the fibonacci sequence
func (svc *FibonacciService) GetCurrent() (int, error) {
	logrus.Info("getting current from db")
	// variable to store query results
	f := types.Fibonacci{}

	svc.client.Database.First(&f, "id = 1")

	// send query to the database and store result in variable
	//db := svc.client.Database.Table("fibonaccis").Where("id = 1").Select("current").Scan(f)
	//

	return f.Current, nil
}

//SetCurrent sets the current value of the fibonacci sequence
func (svc *FibonacciService) SetCurrent(newCurrent int) error {
	f := types.Fibonacci{}

	svc.client.Database.First(&f, "id = 1")

	f.Current = newCurrent

	svc.client.Database.Save(&f)

	return nil
}

//GetPrevious gets the previous value of the fibonacci sequence
func (svc *FibonacciService) GetPrevious() (int, error) {
	// variable to store query results
	f := new(types.Fibonacci)

	// send query to the database and store result in variable
	db := svc.client.Database.Table("fibonaccis").Select("previous").Where("id = 1").Scan(f)

	if db.Error != nil {
		return -1, db.Error
	}

	return f.Previous, nil
}

//SetPrevious sets the previous value of the fibonacci sequence
func (svc *FibonacciService) SetPrevious(newPrev int) error {
	f := types.Fibonacci{}

	svc.client.Database.First(&f, "id = 1")

	f.Previous = newPrev

	svc.client.Database.Save(&f)

	return nil
}
