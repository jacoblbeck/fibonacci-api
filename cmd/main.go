package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "server"
	app.HelpName = "server"
	app.Authors = []*cli.Author{
		{
			Name:  "Jacob Beck",
			Email: "jacobbeck2112@gmail.com",
		},
	}

	app.Action = run
	app.Compiled = time.Now()

	// Server flags
	app.Flags = append(app.Flags, serverFlags...)

	// set logrus to log in JSON format
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

}
