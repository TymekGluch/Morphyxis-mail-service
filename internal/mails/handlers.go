package mails

import (
	"github.com/gofiber/fiber/v3"

	mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"
)

type handlers struct {
	*mailcowIntegration.SmtpMailClient
}

func healthCheckHandler(ctx fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (handlers *handlers) SendAccountConfirmationEmail(ctx fiber.Ctx) error {
	var input SendAccountConfirmationEmailInput

	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := (&input).Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := handlers.SmtpMailClient.SendAccountConfirmationEmail(input.SendAccountConfirmationEmailInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (handlers *handlers) SendAccountVerifiedEmail(ctx fiber.Ctx) error {
	var input SendAccountVerifiedEmailInput

	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := (&input).Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := handlers.SmtpMailClient.SendAccountVerifiedEmail(input.SendAccountVerifiedEmailInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (handlers *handlers) SendPasswordWasChangedEmail(ctx fiber.Ctx) error {
	var input SendPasswordWasChangedEmailInput

	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := (&input).Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := handlers.SmtpMailClient.SendPasswordWasChangedEmail(input.SendPasswordWasChangedEmailInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (handlers *handlers) SendDeletedAccountEmail(ctx fiber.Ctx) error {
	var input SendDeletedAccountEmailInput

	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := (&input).Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := handlers.SmtpMailClient.SendDeletedAccountEmail(input.SendDeletedAccountEmailInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
