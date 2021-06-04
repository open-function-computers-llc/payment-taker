package server

import (
	"testing"
)

func Test_weCanValidateAServerStruct(t *testing.T) {
	s := Server{}
	s.configuration.port = "9000"
	s.configuration.customerName = "ofc"
	s.configuration.stripePublicKey = "i'm required"
	s.configuration.stripePrivateKey = "i'm required"
	s.configuration.smtpHost = "localhost"
	s.configuration.smtpUser = "username"
	s.configuration.smtpPass = "password"
	s.configuration.smtpPort = 1020
	s.configuration.emailFrom = "bot@ofco.email"
	err := validateServer(s)

	if err != nil {
		t.Error("Validation error thrown: " + err.Error())
	}
}

func Test_weGetTheExpectedErrorMessagesWhenThingsAreWrong(t *testing.T) {
	s := Server{}
	err := validateServer(s)
	if err.Error() != "Web port is required" {
		t.Error("expected error 'Web port is required', got: " + err.Error())
	}
	s.configuration.port = "123"

	err = validateServer(s)
	if err.Error() != "Customer name is required" {
		t.Error("expected error 'Customer name is required', got: " + err.Error())
	}
	s.configuration.customerName = "whatever"

	err = validateServer(s)
	if err.Error() != "Stripe public key is not set correctly" {
		t.Error("expected error 'Stripe public key is not set correctly', got: " + err.Error())
	}
	s.configuration.stripePublicKey = "whatever"

	err = validateServer(s)
	if err.Error() != "Stripe private key is not set correctly" {
		t.Error("expected error 'Stripe private key is not set correctly', got: " + err.Error())
	}
	s.configuration.stripePrivateKey = "whatever"

	err = validateServer(s)
	if err.Error() != "SMTP Host is required" {
		t.Error("expected error 'SMTP Host is required', got: " + err.Error())
	}
	s.configuration.smtpHost = "localhost"

	err = validateServer(s)
	if err.Error() != "SMTP User is required" {
		t.Error("expected error 'SMTP User is required', got: " + err.Error())
	}
	s.configuration.smtpUser = "user"

	err = validateServer(s)
	if err.Error() != "SMTP Password is required" {
		t.Error("expected error 'SMTP Password is required', got: " + err.Error())
	}
	s.configuration.smtpPass = "super secret"

	err = validateServer(s)
	if err.Error() != "SMTP Port is required" {
		t.Error("expected error 'SMTP Port is required', got: " + err.Error())
	}
	s.configuration.smtpPort = 1234

	err = validateServer(s)
	if err.Error() != "Email From is required. This field is the automated email messages email address." {
		t.Error("expected error 'Email From is required. This field is the automated email messages email address.', got: " + err.Error())
	}
	s.configuration.emailFrom = "bot@email.com"

	err = validateServer(s)
	if err != nil {
		t.Error("All error cases should be satisfied. Got an additional error: " + err.Error())
	}
}
