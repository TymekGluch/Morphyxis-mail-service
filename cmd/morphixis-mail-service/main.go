package main

import (
	apiDocs "Morphyxis-mail-service/internal/api-docs"
	mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"
	mails "Morphyxis-mail-service/internal/mails"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

const (
	defaultPort = "60"

	serviceIsRunningLog            = "Morphyxis-mail-service is running... \n"
	failedToCreateMailcowClientLog = "Failed to create Mailcow client: %v\n"
)

// @title Morphyxis-mail-service API
// @version 1.0.2
// @description Morphyxis-mail-service API documentation
// @BasePath /api
// @schemes https http
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

	apiDocsService, err := apiDocs.NewService(port)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	apiGroup := app.Group("/api")

	apiDocs.RegisterRoutes(app, apiDocsService)

	mails.RegisterRoutes(apiGroup, &mailcowClient)

	app.Listen(":" + port)
}
