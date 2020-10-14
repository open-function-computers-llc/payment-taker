package server

import (
	"os"

	"github.com/joho/godotenv"
)

func (s *Server) processConfiguration() {
	err := godotenv.Load() // load the env file from the current directory or the parent
	if err != nil {
		s.logError("Couldn't load ENV. Please check your .env file and try again. Continuing with standard ENV loading")
	}
	s.configuration.port = os.Getenv("PORT")
	s.configuration.customerName = os.Getenv("CUSTOMERNAME")
	s.configuration.stripePublicKey = os.Getenv("STRIPEPUBLICKEY")
	s.configuration.stripePrivateKey = os.Getenv("STRIPEPRIVATEKEY")
}
