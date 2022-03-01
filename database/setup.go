package database

import (
	"errors"
	"time"

	"github.com/urfave/cli/v2"
)

type (
	// Setup represents the options that can be set on the database.
	Setup struct {
		Config     string
		Connection *Connection
		Driver     string
	}

	// Connection represents the options that can be set to control the
	// connection settings between the application and the database.
	Connection struct {
		Life time.Duration
		Idle int
		Open int
	}
)

var (
	// ErrInvalidDatabaseDriver defines the error type when the
	// Driver provided to the client is unsupported.
	ErrInvalidDatabaseDriver = errors.New("invalid database driver")

	// Flags represents the required database settings on the cli.
	//nolint
	Flags = []cli.Flag{
		&cli.StringFlag{
			EnvVars: []string{"SERVER_DATABASE_CONFIG", "DATABASE_CONFIG"},
			Name:    "database.config",
			Usage:   "sets the configuration string to be used for the database",
		},
		&cli.IntFlag{
			EnvVars: []string{"SERVER_DATABASE_CONNECTION_OPEN", "DATABASE_CONNECTION_OPEN"},
			Name:    "database.connection.open",
			Usage:   "sets the number of open connections to the database",
			Value:   0,
		},
		&cli.IntFlag{
			EnvVars: []string{"SERVER_DATABASE_CONNECTION_IDLE", "DATABASE_CONNECTION_IDLE"},
			Name:    "database.connection.idle",
			Usage:   "sets the number of idle connections to the database",
			Value:   2,
		},
		&cli.DurationFlag{
			EnvVars: []string{"SERVER_DATABASE_CONNECTION_LIFE", "DATABASE_CONNECTION_LIFE"},
			Name:    "database.connection.life",
			Usage:   "sets the amount of time a connection may be reused for the database",
			Value:   30 * time.Minute,
		},
	}
)
