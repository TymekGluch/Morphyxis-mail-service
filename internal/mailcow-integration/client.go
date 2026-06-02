package mailcowIntegration

import (
	"Morphyxis-mail-service/internal/templates"
	templatesFiles "Morphyxis-mail-service/internal/templates/files"
	"bytes"
	"context"
	"net/smtp"
)

func NewClient(ctx context.Context) (SmtpMailClient, error) {
	config, err := initConfig()
	if err != nil {
		return SmtpMailClient{}, err
	}

	auth := smtp.PlainAuth(
		"",
		config.User,
		config.Password,
		config.Host,
	)

	return SmtpMailClient{ctx: ctx, smtpConfig: config, smtpAuth: auth}, nil
}

func (client SmtpMailClient) SendAccountConfirmationEmail(input SendAccountConfirmationEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	var body bytes.Buffer

	template := templatesFiles.AccountConfirmation(templates.AccountConfirmationInput{
		Name:                input.Name,
		VerificationCode:    input.VerificationCode,
		AccountDeletionDate: input.AccountDeletionDate,
	})

	if err := template.Render(client.ctx, &body); err != nil {
		return err
	}

	message := createSmtpMessage(
		client.User,
		input.To,
		input.Subject,
		body.String(),
	)

	err := smtp.SendMail(
		client.Host+":"+client.Port,
		client.smtpAuth,
		client.User,
		[]string{input.To},
		message,
	)

	if err != nil {
		return err
	}

	return nil
}

func (client SmtpMailClient) SendAccountVerifiedEmail(input SendAccountVerifiedEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	var body bytes.Buffer

	template := templatesFiles.AccountVerified(templates.AccountVerifiedInput{
		Name: input.Name,
	})

	if err := template.Render(client.ctx, &body); err != nil {
		return err
	}

	message := createSmtpMessage(
		client.User,
		input.To,
		input.Subject,
		body.String(),
	)

	err := smtp.SendMail(
		client.Host+":"+client.Port,
		client.smtpAuth,
		client.User,
		[]string{input.To},
		message,
	)

	if err != nil {
		return err
	}

	return nil
}

func (client SmtpMailClient) SendPasswordWasChangedEmail(input SendPasswordWasChangedEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	var body bytes.Buffer

	template := templatesFiles.PasswordWasChanged(templates.PasswordWasChangedInput{
		Name:          input.Name,
		DateOfRequest: input.DateOfRequest,
	})

	if err := template.Render(client.ctx, &body); err != nil {
		return err
	}

	message := createSmtpMessage(
		client.User,
		input.To,
		input.Subject,
		body.String(),
	)

	err := smtp.SendMail(
		client.Host+":"+client.Port,
		client.smtpAuth,
		client.User,
		[]string{input.To},
		message,
	)

	if err != nil {
		return err
	}

	return nil
}

func (client SmtpMailClient) SendDeletedAccountEmail(input SendDeletedAccountEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	var body bytes.Buffer

	template := templatesFiles.DeletedAccount(templates.DeletedAccountInput{
		Name:   input.Name,
		Reason: input.Reason,
	})

	if err := template.Render(client.ctx, &body); err != nil {
		return err
	}

	message := createSmtpMessage(
		client.User,
		input.To,
		input.Subject,
		body.String(),
	)

	err := smtp.SendMail(
		client.Host+":"+client.Port,
		client.smtpAuth,
		client.User,
		[]string{input.To},
		message,
	)

	if err != nil {
		return err
	}

	return nil
}
