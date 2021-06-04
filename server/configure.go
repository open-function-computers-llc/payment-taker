package server

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func (s *Server) processConfiguration() error {
	err := godotenv.Load() // load the env file from the current directory or the parent
	if err != nil {
		s.logError("Couldn't load ENV. Please check your .env file and try again. Continuing with standard ENV loading")
		return err
	}
	s.configuration.port = os.Getenv("PORT")
	s.configuration.customerName = os.Getenv("CUSTOMERNAME")
	s.configuration.stripePublicKey = os.Getenv("STRIPEPUBLICKEY")
	s.configuration.stripePrivateKey = os.Getenv("STRIPEPRIVATEKEY")

	s.configuration.smtpHost = os.Getenv("SMTP_HOST")
	s.configuration.smtpUser = os.Getenv("SMTP_USER")
	s.configuration.smtpPass = os.Getenv("SMTP_PASSWORD")
	smtpPortString := os.Getenv("SMTP_PORT")
	s.configuration.smtpPort, err = strconv.Atoi(smtpPortString)
	if err != nil {
		s.logError("Your port number is not set correctly in your ENV. Check the value for SMTP_PORT")
		return err
	}

	s.configuration.emailFrom = os.Getenv("EMAIL_FROM")
	s.configuration.systemNotificationTo = os.Getenv("EMAIL_TO")

	return nil
}
