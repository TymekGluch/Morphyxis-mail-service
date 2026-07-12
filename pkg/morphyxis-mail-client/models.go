package morphyxisMailClient

import (
	"net/http"
	"time"
)

type Config struct {
	BaseURL string
	Timeout time.Duration
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

type SendAccountConfirmationEmailInput struct {
	To                  string    `json:"to"`
	Subject             string    `json:"subject"`
	Name                string    `json:"name"`
	VerificationCode    string    `json:"verification_code"`
	AccountDeletionDate time.Time `json:"account_deletion_date"`
}

type SendAccountVerifiedEmailInput struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
}

type SendPasswordWasChangedEmailInput struct {
	To            string    `json:"to"`
	Subject       string    `json:"subject"`
	Name          string    `json:"name"`
	DateOfRequest time.Time `json:"date_of_request"`
}

type SendDeletedAccountEmailInput struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
	Reason  string `json:"reason"`
}
