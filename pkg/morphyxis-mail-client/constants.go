package morphyxisMailClient

import "time"

const (
	defaultTimeout = 10 * time.Second

	healthCheckPath         = "/api/health-check"
	accountConfirmationPath = "/api/emails/account-confirmation"
	accountVerifiedPath     = "/api/emails/account-verified"
	passwordChangedPath     = "/api/emails/password-was-changed"
	deletedAccountPath      = "/api/emails/deleted-account"
)
