package mailcowIntegration

import (
	"fmt"
	"os"
)

type smtpConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Domain   string
}

const (
	hostMissingError     = "MAILCOW_SMTP_HOST environment variable is not set"
	portMissingError     = "MAILCOW_SMTP_PORT environment variable is not set"
	userMissingError     = "MAILCOW_USER environment variable is not set"
	passwordMissingError = "MAILCOW_MAILBOX_PASSWORD environment variable is not set"
)

func initConfig() (smtpConfig, error) {
	host := os.Getenv("MAILCOW_SMTP_HOST")
	if host == "" {
		return smtpConfig{}, fmt.Errorf(hostMissingError)
	}

	domain := os.Getenv("MAILCOW_SMTP_DOMAIN")
	if domain == "" {
		return smtpConfig{}, fmt.Errorf("MAILCOW_SMTP_DOMAIN environment variable is not set")
	}

	port := os.Getenv("MAILCOW_SMTP_PORT")
	if port == "" {
		return smtpConfig{}, fmt.Errorf(portMissingError)
	}

	user := os.Getenv("MAILCOW_USER")
	if user == "" {
		return smtpConfig{}, fmt.Errorf(userMissingError)
	}

	password := os.Getenv("MAILCOW_MAILBOX_PASSWORD")
	if password == "" {
		return smtpConfig{}, fmt.Errorf(passwordMissingError)
	}

	return smtpConfig{
		Host:     host,
		Domain:   domain,
		Port:     port,
		User:     user,
		Password: password,
	}, nil
}
