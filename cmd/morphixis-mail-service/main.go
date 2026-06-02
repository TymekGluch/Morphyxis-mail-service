package main

import (
	mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"
	"context"
	"fmt"
)

func main() {
	fmt.Print("Morphyxis-mail-service is running... \n")

	ctx := context.Background()
	mailClient, err := mailcowIntegration.NewClient(ctx)
	if err != nil {
		fmt.Println("Error creating mail client:", err)
		return
	}

	err = mailClient.SendDeletedAccountEmail(mailcowIntegration.SendDeletedAccountEmailInput{
		To:      "tymoteuszgluch848@gmail.com",
		Subject: "Account Deleted",
		Name:    "test user",
		Reason:  "User requested account deletion",
	})
	if err != nil {
		fmt.Println("Error sending deleted account email:", err)
		return
	} else {
		fmt.Println("Deleted account email sent successfully!")
	}

}
