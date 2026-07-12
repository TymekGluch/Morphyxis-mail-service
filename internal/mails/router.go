package mails

import (
	mailcowIntegration "github.com/TymekGluch/Morphyxis-mail-service/internal/mailcow-integration"
	"github.com/TymekGluch/Morphyxis-mail-service/internal/timeouts"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app fiber.Router, mailcowClient *mailcowIntegration.SmtpMailClient) {
	handlers := &handlers{mailcowClient}

	app.Get("/health-check", healthCheckHandler, timeouts.WithDefaultTimeoutHandler())

	emailGroup := app.Group("/emails", timeouts.WithDefaultTimeoutHandler())

	emailGroup.Post("/account-confirmation", handlers.SendAccountConfirmationEmail)
	emailGroup.Post("/account-verified", handlers.SendAccountVerifiedEmail)
	emailGroup.Post("/password-was-changed", handlers.SendPasswordWasChangedEmail)
	emailGroup.Post("/deleted-account", handlers.SendDeletedAccountEmail)
}
