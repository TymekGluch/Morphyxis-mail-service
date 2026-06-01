package mailcowIntegration

import (
	"context"
	"fmt"
	"net/smtp"
	"time"
)

type SendEmailInput struct {
	To      string
	Subject string
}

func (input SendEmailInput) Validate() error {
	if input.To == "" {
		return fmt.Errorf(mailcowInputToValidationError)
	}

	if input.Subject == "" {
		return fmt.Errorf(mailcowInputSubjectValidationError)
	}

	return nil
}

type SmtpMailClient struct {
	ctx context.Context
	smtpConfig
	smtpAuth smtp.Auth
}

type SendAccountConfirmationEmailInput struct {
	To                  string
	Subject             string
	Name                string
	VerificationCode    string
	AccountDeletionDate time.Time
}

func (input SendAccountConfirmationEmailInput) Validate() error {
	if To := input.To; To == "" {
		return fmt.Errorf(mailcowInputToValidationError)
	}

	if Subject := input.Subject; Subject == "" {
		return fmt.Errorf(mailcowInputSubjectValidationError)
	}

	return nil
}
