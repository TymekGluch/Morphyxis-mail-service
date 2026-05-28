package templates

import "time"

type AccountConfirmationInput struct {
	Name                string
	VerificationCode    string
	AccountDeletionDate time.Time
}

func (accountConfirmationInput AccountConfirmationInput) AccountDeletionDateFormatted() string {
	return accountConfirmationInput.AccountDeletionDate.Format("02.01.2006 15:04")
}

type DeletedAccountInput struct {
	Name   string
	Reason string
}

type AccountVerifiedInput struct {
	Name string
}

type PasswordWasChangedInput struct {
	Name          string
	DateOfRequest time.Time
}

func (passwordWasChangedInput PasswordWasChangedInput) DateOfRequestFormatted() string {
	return passwordWasChangedInput.DateOfRequest.Format("02.01.2006 15:04")
}
