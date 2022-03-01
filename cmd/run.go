package main

import "github.com/urfave/cli/v2"

func run(c *cli.Context) error {

	s := server{
		Addr: c.String("server.addr"),
		Port: c.String("server.port"),
	}

	s.configure()

	return s.start()
}
