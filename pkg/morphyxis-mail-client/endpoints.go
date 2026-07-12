package morphyxisMailClient

import (
	"context"
	"fmt"
	"net/http"
)

func (client *client) HealthCheck(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.baseURL+healthCheckPath, nil)
	if err != nil {
		return err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("morphyxis-mail-client: health check returned status %d", resp.StatusCode)
	}

	return nil
}

func (client *client) SendAccountConfirmationEmail(ctx context.Context, input SendAccountConfirmationEmailInput) error {
	return client.post(ctx, accountConfirmationPath, input)
}

func (client *client) SendAccountVerifiedEmail(ctx context.Context, input SendAccountVerifiedEmailInput) error {
	return client.post(ctx, accountVerifiedPath, input)
}

func (client *client) SendPasswordWasChangedEmail(ctx context.Context, input SendPasswordWasChangedEmailInput) error {
	return client.post(ctx, passwordChangedPath, input)
}

func (client *client) SendDeletedAccountEmail(ctx context.Context, input SendDeletedAccountEmailInput) error {
	return client.post(ctx, deletedAccountPath, input)
}
