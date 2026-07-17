# Changelog

## [1.1.3] - 2026-07-17

### Fixed

- Add correct importing module in `mail-template-sandbox`

## [1.1.1] - 2026-07-17

### Fixed

- Fixed incorrect module name in all app

## [1.1.0] - 2026-07-12

### Added

- Added go client package, which support `morphyxis-mail-service` API.

```go
package main

import (
  "context"
  "log"
  "time"

  mailClient "github.com/morphyxis/morphyxis-mail-service/pkg/morphyxis-mail-client"
)

func main() {
  mc, err := mailClient.New(mailClient.Config{
    BaseURL: "https://mail.example.com",
  })
  if err != nil {
    log.Fatal(err)
  }

  ctx := context.Background()

  err = mc.SendAccountConfirmationEmail(ctx, mailClient.SendAccountConfirmationEmailInput{
    To:                  "user@example.com",
    Subject:             "Confirm your account",
    Name:                "Jan Kowalski",
    VerificationCode:    "ABC123",
    AccountDeletionDate: time.Now().Add(30 * 24 * time.Hour),
  })
  if err != nil {
    log.Fatal(err)
  }
}
```

## [1.0.2] - 2026-07-12

- Added configurable independent TLS Server name and Mailcow port

## [1.0.1] - 2026-07-12

### Added

- Added configurable SMTP TLS Server Name.

### Fixed

- Fixed TLS certificate validation when using Mailcow's internal Postfix service.

## [1.0.0] - 2026-07-01

### Added

- Initial release.
