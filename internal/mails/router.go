package mails

import (
	mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app fiber.Router, mailcowClient *mailcowIntegration.SmtpMailClient) {
	handlers := &handlers{mailcowClient}

	emailGroup := app.Group("/emails")

	emailGroup.Get("/health-check", healthCheckHandler)
	emailGroup.Post("/account-confirmation", handlers.SendAccountConfirmationEmail)
	emailGroup.Post("/account-verified", handlers.SendAccountVerifiedEmail)
	emailGroup.Post("/password-was-changed", handlers.SendPasswordWasChangedEmail)
	emailGroup.Post("/deleted-account", handlers.SendDeletedAccountEmail)
}
