package mails

import (
	"github.com/gofiber/fiber/v3"

	mailcowIntegration "github.com/TymekGluch/Morphyxis-mail-service/internal/mailcow-integration"
)

type handlers struct {
	*mailcowIntegration.SmtpMailClient
}

// HealthCheckHandler godoc
// @Summary Health Check Handler
// @Description Health Check Handler
// @Tags Health Health Check
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 408 {object} error "Request Timeout"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/health-check [get]
func healthCheckHandler(ctx fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

// SendAccountConfirmationEmail godoc
// @Summary Send Account Confirmation Email
// @Description Sends an account confirmation email with a verification code
// @Tags Emails
// @Accept json
// @Produce json
// @Param input body SendAccountConfirmationEmailInput true "Account Confirmation Email Input"
// @Success 200 {string} string "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 408 {object} error "Request Timeout"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/emails/account-confirmation [post]
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

// SendAccountVerifiedEmail godoc
// @Summary Send Account Verified Email
// @Description Sends an email notifying the user that their account has been verified
// @Tags Emails
// @Accept json
// @Produce json
// @Param input body SendAccountVerifiedEmailInput true "Account Verified Email Input"
// @Success 200 {string} string "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 408 {object} error "Request Timeout"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/emails/account-verified [post]
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

// SendPasswordWasChangedEmail godoc
// @Summary Send Password Was Changed Email
// @Description Sends an email notifying the user that their password has been changed
// @Tags Emails
// @Accept json
// @Produce json
// @Param input body SendPasswordWasChangedEmailInput true "Password Was Changed Email Input"
// @Success 200 {string} string "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 408 {object} error "Request Timeout"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/emails/password-was-changed [post]
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

// SendDeletedAccountEmail godoc
// @Summary Send Deleted Account Email
// @Description Sends an email notifying the user that their account has been deleted
// @Tags Emails
// @Accept json
// @Produce json
// @Param input body SendDeletedAccountEmailInput true "Deleted Account Email Input"
// @Success 200 {string} string "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 408 {object} error "Request Timeout"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/emails/deleted-account [post]
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
