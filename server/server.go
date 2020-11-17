package server

import (
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

// Server is a struct that controls all the input/ouput through http
type Server struct {
	router        *mux.Router
	logger        *logrus.Logger
	staticAssets  *rice.Box
	views         *rice.Box
	configuration struct {
		port             string
		customerName     string
		stripePublicKey  string
		stripePrivateKey string
	}
}

// Create return a new instance of a server booted up and ready to go
func Create() Server {
	s := Server{}

	// set up logging
	s.logger = logrus.New()
	s.logger.SetOutput(os.Stdout)

	// set up static asset and tpl riceboxes
	s.staticAssets = rice.MustFindBox("../static")
	s.views = rice.MustFindBox("../views")

	// set up routing
	s.router = mux.NewRouter()
	s.setRoutes()

	// set app configuration from ENV
	s.processConfiguration()

	return s
}

// Serve will start up the http server
func (s *Server) Serve() error {
	s.log("Starting server on port " + s.configuration.port)

	// CORS stuff
	rules := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})
	corsMux := rules.Handler(s.router)
	return http.ListenAndServe(":"+s.configuration.port, corsMux)
}

func (s *Server) log(messages ...interface{}) {
	s.logger.Info(messages...)
}

func (s *Server) logError(message interface{}) {
	s.logger.Error(message)
}
