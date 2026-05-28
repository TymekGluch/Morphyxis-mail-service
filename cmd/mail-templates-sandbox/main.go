package main

import (
	"Morphyxis-mail-service/internal/templates"
	templatesFiles "Morphyxis-mail-service/internal/templates/files"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx fiber.Ctx) error {
		templates := []string{
			"confirmationAccount",
			"deletedAccount",
			"accountVerified",
			"passwordWasChanged",
		}

		return ctx.Render("views/index.html", fiber.Map{
			"Templates": templates,
		})
	})

	app.Get("/templates/confirmationAccount", templ.Handler(templatesFiles.AccountConfirmation(templates.AccountConfirmationInput{
		Name:                "Park Dildos Smith",
		VerificationCode:    "123456d",
		AccountDeletionDate: time.Now().Add(24 * 7 * time.Hour),
	})))

	app.Get("/templates/deletedAccount", templ.Handler(templatesFiles.DeletedAccount(templates.DeletedAccountInput{
		Name:   "Park Dildos Smith",
		Reason: "You have not confirmed your account within 7 days.",
	})))

	app.Get("/templates/accountVerified", templ.Handler(templatesFiles.AccountVerified(templates.AccountVerifiedInput{
		Name: "Park Dildos Smith ",
	})))

	app.Get("/templates/passwordWasChanged", templ.Handler(templatesFiles.PasswordWasChanged(templates.PasswordWasChangedInput{
		Name:          "Park Dildos Smith ",
		DateOfRequest: time.Now(),
	})))

	app.Listen(":50")
}
