package morphyxisMailClient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailServiceClient interface {
	HealthCheck(ctx context.Context) error
	SendAccountConfirmationEmail(ctx context.Context, input SendAccountConfirmationEmailInput) error
	SendAccountVerifiedEmail(ctx context.Context, input SendAccountVerifiedEmailInput) error
	SendPasswordWasChangedEmail(ctx context.Context, input SendPasswordWasChangedEmailInput) error
	SendDeletedAccountEmail(ctx context.Context, input SendDeletedAccountEmailInput) error
}

func New(cfg Config) (MailServiceClient, error) {
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("morphyxis-mail-client: BaseURL is required")
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}

	return &client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (client *client) post(ctx context.Context, path string, body any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, client.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("morphyxis-mail-client: %s returned status %d", path, response.StatusCode)
	}

	return nil
}
