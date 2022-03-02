package main

import (
	"github.com/jacoblbeck/fibonacci-api/database"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	// set log level for logrus
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "i", "info", "Info", "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	}

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
