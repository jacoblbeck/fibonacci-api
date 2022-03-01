package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jacoblbeck/fibonacci-api/database"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/tomb.v2"
)

// Server represents the information that can be provided
// to run a Marina server.
type server struct {
	Addr     string
	Database *database.Setup
	Port     string
	Router   *mux.Router
}

// serverFlags represents the settings that can be passed from the cli.
var serverFlags = []cli.Flag{
	&cli.StringFlag{
		EnvVars: []string{"SERVER_ADDR", "ADDR"},
		Name:    "server.addr",
		Usage:   "server address as a fully qualified url (<scheme>://<host>)",
	},
	&cli.StringFlag{
		EnvVars: []string{"SERVER_PORT", "PORT"},
		Name:    "server.port",
		Usage:   "API port to listen on",
		Value:   ":8080",
	},
	&cli.StringFlag{
		EnvVars: []string{"SERVER_LOG_LEVEL", "LOG_LEVEL"},
		Name:    "log.level",
		Usage:   "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
		Value:   "info",
	},
}

// configure is the function thats creates and configures
// all the required dependencies for the server
func (s *server) configure() error {
	// create a new database configuration
	database, err := database.New(s.Database)
	if err != nil {
		return err
	}

	// overwrite the router with the required middleware
	s.Router = router(database.Middleware)

	return nil
}

// start is the function thats to listen and serve
// traffic for web and API requests.
func (s *server) start() error {
	var tomb tomb.Tomb

	// start http server
	tomb.Go(func() error {
		srv := &http.Server{Addr: s.Port, Handler: s.Router}

		go func() {
			logrus.Info("Starting HTTP server...")
			err := srv.ListenAndServe()
			if err != nil {
				tomb.Kill(err)
			}
		}()

		for {
			select {
			case <-tomb.Dying():
				logrus.Info("Stopping HTTP server...")
				return srv.Shutdown(context.Background())
			}
		}
	})

	// Wait for stuff and watch for errors
	err := tomb.Wait()
	if err != nil {
		return err
	}

	return tomb.Err()
}

// Validate verifies the necessary fields for the
// provided configuration are populated correctly.
func (s *server) validate() error {
	logrus.Debug("validating server configuration")

	// validate the server setup

	if len(s.Addr) == 0 {
		return fmt.Errorf("server.addr (SERVER_ADDR or ADDR) flag is not properly configured")
	}

	if !strings.Contains(s.Addr, "://") {
		return fmt.Errorf("server.addr (SERVER_ADDR or ADDR) flag must be <scheme>://<hostname> format")
	}

	if strings.HasSuffix(s.Addr, "/") {
		return fmt.Errorf("server.addr (SERVER_ADDR or ADDR) flag must not have trailing slash")
	}

	if len(s.Port) == 0 {
		return fmt.Errorf("server.port (SERVER_PORT or PORT) flag is not properly configured")
	}

	return nil
}
