package database

import (
	"time"

	"github.com/jacoblbeck/fibonacci-api/types"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type service struct {
	client *Client
}

// Client represents a struct to hold database setup.
type Client struct {
	Database  *gorm.DB
	Fibonacci *FibonacciService
}

// New returns a Database implementation that
// integrates with a supported database instance.
func New(s *Setup) (*Client, error) {
	// create the database client
	retries := 0

	db, err := gorm.Open(postgres.Open(s.Config), &gorm.Config{})

	//retry for database connection if server starts before db is ready
	for err != nil {
		if retries < 2 {
			retries++
			time.Sleep(5 * time.Second)
			db, err = gorm.Open(postgres.Open(s.Config), &gorm.Config{
				SkipDefaultTransaction: true,
				PrepareStmt:            true,
			})
			continue
		}
		logrus.Error(err)
	}

	err = db.AutoMigrate(types.Fibonacci{})
	if err != nil {
		return nil, err
	}

	var exists bool
	err = db.Model(types.Fibonacci{}).
		Select("count(*) > 0").
		Where("current::text LIKE ?", "%").
		Find(&exists).
		Error

	if err != nil {
		return nil, err
	}

	if !exists {
		db.Create(&types.Fibonacci{Current: 0, Previous: 0})
	}

	database, err := db.DB()
	if err != nil {
		return nil, err
	}

	// apply extra database configuration
	database.SetConnMaxLifetime(s.Connection.Life)
	database.SetMaxIdleConns(s.Connection.Idle)
	database.SetMaxOpenConns(s.Connection.Open)

	// create the client object
	client := &Client{
		Database: db,
	}

	client.Fibonacci = &FibonacciService{client: client}

	return client, nil
}
