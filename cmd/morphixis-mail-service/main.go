package main

import (
	mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"
	mails "Morphyxis-mail-service/internal/mails"
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

const (
	defaultPort = "60"

	serviceIsRunningLog            = "Morphyxis-mail-service is running... \n"
	failedToCreateMailcowClientLog = "Failed to create Mailcow client: %v\n"
)

func main() {
	port := os.Getenv("MORPHYXIS_MAIL_SERVICE_PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Print(serviceIsRunningLog)

	ctx := context.Background()

	mailcowClient, err := mailcowIntegration.NewClient(ctx)
	if err != nil {
		fmt.Printf(failedToCreateMailcowClientLog, err)
		return
	}

	app := fiber.New()
	apiGroup := app.Group("/api")

	mails.RegisterRoutes(apiGroup, &mailcowClient)

	app.Listen(":" + port)
}
