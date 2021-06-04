package server

import "errors"

func validateServer(s Server) error {
	if s.configuration.port == "" {
		return errors.New("Web port is required")
	}
	if s.configuration.customerName == "" {
		return errors.New("Customer name is required")
	}
	if s.configuration.stripePublicKey == "" {
		return errors.New("Stripe public key is not set correctly")
	}
	if s.configuration.stripePrivateKey == "" {
		return errors.New("Stripe private key is not set correctly")
	}

	// email sending
	if s.configuration.smtpHost == "" {
		return errors.New("SMTP Host is required")
	}
	if s.configuration.smtpUser == "" {
		return errors.New("SMTP User is required")
	}
	if s.configuration.smtpPass == "" {
		return errors.New("SMTP Password is required")
	}
	if s.configuration.smtpPort == 0 {
		return errors.New("SMTP Port is required")
	}
	if s.configuration.emailFrom == "" {
		return errors.New("Email From is required. This field is the automated email messages email address.")
	}
	return nil
}
