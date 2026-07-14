package main

import (
	"context"
	"fmt"
	"log"
	"os"

	apiDocs "github.com/TymekGluch/Morphyxis-mail-service/internal/api-docs"
	mailcowIntegration "github.com/TymekGluch/Morphyxis-mail-service/internal/mailcow-integration"
	mails "github.com/TymekGluch/Morphyxis-mail-service/internal/mails"

	"github.com/gofiber/fiber/v3"
)

const (
	defaultPort = "60"

	serviceIsRunningLog            = "Morphyxis-mail-service is running... \n"
	failedToCreateMailcowClientLog = "Failed to create Mailcow client: %v\n"
)

// @title Morphyxis-mail-service API
// @version 1.1.1
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
