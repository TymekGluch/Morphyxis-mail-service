package mailcowIntegration

import (
	"Morphyxis-mail-service/internal/templates"
	templatesFiles "Morphyxis-mail-service/internal/templates/files"
	"bytes"
	"context"
	"crypto/tls"
	"strconv"

	"github.com/wneessen/go-mail"
)

func NewClient(ctx context.Context) (SmtpMailClient, error) {
	config, err := initConfig()
	if err != nil {
		return SmtpMailClient{}, err
	}

	return SmtpMailClient{ctx: ctx, smtpConfig: config}, nil
}

func (client SmtpMailClient) newMailClient() (*mail.Client, error) {
	port, err := strconv.Atoi(client.Port)
	if err != nil {
		return nil, err
	}

	return mail.NewClient(
		client.Host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(client.User),
		mail.WithPassword(client.Password),
		mail.WithTLSConfig(&tls.Config{ServerName: client.Domain}),
	)
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

	msg := mail.NewMsg()
	if err := msg.From(client.User); err != nil {
		return err
	}
	if err := msg.To(input.To); err != nil {
		return err
	}
	msg.Subject(input.Subject)
	msg.SetBodyString(mail.TypeTextHTML, body.String())

	mailClient, err := client.newMailClient()
	if err != nil {
		return err
	}

	return mailClient.DialAndSend(msg)
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

	msg := mail.NewMsg()
	if err := msg.From(client.User); err != nil {
		return err
	}
	if err := msg.To(input.To); err != nil {
		return err
	}
	msg.Subject(input.Subject)
	msg.SetBodyString(mail.TypeTextHTML, body.String())

	mailClient, err := client.newMailClient()
	if err != nil {
		return err
	}

	return mailClient.DialAndSend(msg)
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

	msg := mail.NewMsg()
	if err := msg.From(client.User); err != nil {
		return err
	}
	if err := msg.To(input.To); err != nil {
		return err
	}
	msg.Subject(input.Subject)
	msg.SetBodyString(mail.TypeTextHTML, body.String())

	mailClient, err := client.newMailClient()
	if err != nil {
		return err
	}

	return mailClient.DialAndSend(msg)
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

	msg := mail.NewMsg()
	if err := msg.From(client.User); err != nil {
		return err
	}
	if err := msg.To(input.To); err != nil {
		return err
	}
	msg.Subject(input.Subject)
	msg.SetBodyString(mail.TypeTextHTML, body.String())

	mailClient, err := client.newMailClient()
	if err != nil {
		return err
	}

	return mailClient.DialAndSend(msg)
}
