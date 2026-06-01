package main

import (
	"fmt"
	// mailcowIntegration "Morphyxis-mail-service/internal/mailcow-integration"
	// "context"
	// "time"
)

func main() {
	fmt.Print("Morphyxis-mail-service is running... \n")

	// ctx := context.Background()
	// mailClient, err := mailcowIntegration.NewClient(ctx)
	// if err != nil {
	// 	fmt.Println("Error creating mail client:", err)
	// 	return
	// }

	// err = mailClient.SendAccountConfirmationEmail(mailcowIntegration.SendAccountConfirmationEmailInput{
	// 	To:                  "",
	// 	Subject:             "Account Confirmation",
	// 	Name:                "test user",
	// 	VerificationCode:    "123456d",
	// 	AccountDeletionDate: time.Now().Add(24 * 7 * time.Hour),
	// })
	// if err != nil {
	// 	fmt.Println("Error sending account confirmation email:", err)
	// 	return
	// } else {
	// 	fmt.Println("Account confirmation email sent successfully!")
	// }
}
