package mails

import (
	mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"

	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type SendAccountConfirmationEmailInput struct {
	mailcowIntegration.SendAccountConfirmationEmailInput
}

func (input SendAccountConfirmationEmailInput) Validate() error {
	err := validation.ValidateStruct(
		&input,
		validation.Field(
			&input.To,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&input.Name,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(
			&input.Subject,
			validation.Required,
			validation.Length(3, 140),
		),
		validation.Field(
			&input.VerificationCode,
			validation.Required,
			validation.Length(6, 6),
		),
		validation.Field(
			&input.AccountDeletionDate,
			validation.Required,
		),
	)

	return err
}

type SendAccountVerifiedEmailInput struct {
	mailcowIntegration.SendAccountVerifiedEmailInput
}

func (input SendAccountVerifiedEmailInput) Validate() error {
	err := validation.ValidateStruct(
		&input,
		validation.Field(
			&input.To,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&input.Name,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(
			&input.Subject,
			validation.Required,
			validation.Length(3, 140),
		),
	)

	return err
}

type SendPasswordWasChangedEmailInput struct {
	mailcowIntegration.SendPasswordWasChangedEmailInput
}

func (input SendPasswordWasChangedEmailInput) Validate() error {
	err := validation.ValidateStruct(
		&input,
		validation.Field(
			&input.To,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&input.Name,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(
			&input.Subject,
			validation.Required,
			validation.Length(3, 140),
		),
		validation.Field(
			&input.DateOfRequest,
			validation.Required,
		),
	)

	return err
}

type SendDeletedAccountEmailInput struct {
	mailcowIntegration.SendDeletedAccountEmailInput
}

func (input SendDeletedAccountEmailInput) Validate() error {
	err := validation.ValidateStruct(
		&input,
		validation.Field(
			&input.To,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&input.Name,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(
			&input.Subject,
			validation.Required,
			validation.Length(3, 140),
		),
		validation.Field(
			&input.Reason,
			validation.Required,
			validation.Length(3, 140),
		),
	)

	return err
}
