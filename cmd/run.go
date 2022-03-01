package main

import (
	"github.com/jacoblbeck/fibonacci-api/database"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {

	s := server{
		Addr: c.String("server.addr"),
		Port: c.String("server.port"),
		Database: &database.Setup{
			Config: c.String("database.config"),
			Connection: &database.Connection{
				Life: c.Duration("database.connection.life"),
				Idle: c.Int("database.connection.idle"),
				Open: c.Int("database.connection.open"),
			},
		},
	}

	s.configure()

	return s.start()
}
