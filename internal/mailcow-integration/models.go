package mailcowIntegration

import (
	"context"
	"fmt"
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
}

type SendAccountConfirmationEmailInput struct {
	To                  string    `json:"to"`
	Subject             string    `json:"subject"`
	Name                string    `json:"name"`
	VerificationCode    string    `json:"verification_code"`
	AccountDeletionDate time.Time `json:"account_deletion_date"`
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

type SendAccountVerifiedEmailInput struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
}

func (input SendAccountVerifiedEmailInput) Validate() error {
	if To := input.To; To == "" {
		return fmt.Errorf(mailcowInputToValidationError)
	}

	if Subject := input.Subject; Subject == "" {
		return fmt.Errorf(mailcowInputSubjectValidationError)
	}

	if Name := input.Name; Name == "" {
		return fmt.Errorf(mailcowInputNameValidationError)
	}

	return nil
}

type SendPasswordWasChangedEmailInput struct {
	To            string    `json:"to"`
	Subject       string    `json:"subject"`
	Name          string    `json:"name"`
	DateOfRequest time.Time `json:"date_of_request"`
}

func (input SendPasswordWasChangedEmailInput) Validate() error {
	if To := input.To; To == "" {
		return fmt.Errorf(mailcowInputToValidationError)
	}

	if Subject := input.Subject; Subject == "" {
		return fmt.Errorf(mailcowInputSubjectValidationError)
	}

	if Name := input.Name; Name == "" {
		return fmt.Errorf(mailcowInputNameValidationError)
	}

	return nil
}

type SendDeletedAccountEmailInput struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
	Reason  string `json:"reason"`
}

func (input SendDeletedAccountEmailInput) Validate() error {
	if To := input.To; To == "" {
		return fmt.Errorf(mailcowInputToValidationError)
	}

	if Subject := input.Subject; Subject == "" {
		return fmt.Errorf(mailcowInputSubjectValidationError)
	}

	if Name := input.Name; Name == "" {
		return fmt.Errorf(mailcowInputNameValidationError)
	}

	if Reason := input.Reason; Reason == "" {
		return fmt.Errorf(mailcowInputReasonValidationError)
	}

	return nil
}
