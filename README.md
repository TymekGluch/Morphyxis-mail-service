# MORPHYXIS MAIL SERVICE

Small Open source go lang package to handle Own Mail Server (Mailcow based) with common mail template and handle common requests like POST request to send mails with custom templates

## Current Version

```bash
1.1.0
```

## Get Started

1.  to start working with code, execute command

```bash
make init
```

- to copy scripts to git hooks (pre-push, pre-commit etc.)

## Before installation

This container must be connected to the `mailcowdockerized_mailcow-network` Docker network.
The Mailcow nginx configuration must be extended or overridden to proxy requests to this service.

- WARNING: This service does not implement authorization. I highly recommend securing it with mutual TLS (mTLS).

## GHCR Registry

- `ghcr.io/morphyxis/morphyxis-mail-service` - contains docker images for morphyxis mail service
  - `1.1.0` - latest stable version of the image (tagged with git tag & build manually triggered)
  - `beta` - latest beta version of the image (built on pull request)

to pull the image, execute command

```bash
  docker pull ghcr.io/morphyxis/morphyxis-mail-service:1.1.0
```

to run the image on your server:

```bash
docker run -d \
  -e MAILCOW_MAILBOX_PASSWORD=somePassword \
  -e MAILCOW_SMTP_HOST=postfix-mailcow \
  -e MAILCOW_SMTP_DOMAIN=mail.domain.com \
  -e MAILCOW_SMTP_PORT=453 \
  -e MAILCOW_USER=someUser \
  -e MORPHYXIS_MAIL_SERVICE_PORT=8080 \
  -p 8080:8080 \
  ghcr.io/morphyxis/morphyxis-mail-service:1.1.0
```

or with file `.env` (recommended):

```bash
docker pull ghcr.io/morphyxis/morphyxis-mail-service:1.1.0
docker run -d \
  --env-file .env \
  -p 8080:8080 \
  ghcr.io/morphyxis/morphyxis-mail-service:1.1.0
```

## Project Structure

- `github/workflow` - contains github workflow files for CI/CD
  - `update-image.yml` - pipelines to build and push docker images to ghcr registry
- `cmd` - contains files to start the multiple services
  - `morphixis-mail-service.go` - main file to start the mail service
  - `mail-templates-sandbox` - file to start the mail templates sandbox (FOR DEVELOPMENT PURPOSE ONLY)
- `pkg/` - contains go packages for the service
  - `morphyxis-mail-client` - contains go client to interact with morphyxis mail service
    - `client.go` - contains client to interact with morphyxis mail service
    - `config.go` - contains configuration for the client
    - `constants.go` - contains constants for the client
    - `helpers.go` - contains helper functions for the client
    - `models.go` - contains models for the client
- `internal` - contains internal packages for the service
  - `api-docs` - contains swagger documentation for the service
  - `docs` - genderated swagger documentation for the service
  - `mailcow-integration` - contains integration with Mailcow API (f.e. smtp requests and handler for all mail templates)
    - `client.go` - contains client to interact with mailcow API
    - `config.go` - contains credentials for the mailcow client
    - `constants.go`,
    - `helpers.go` - contains helper functions
    - `models.go` - contains models for mailcow API
  - `mails` - contains module to handle mails
    - `handlers.go` - contains handler to send mails with mailcow API
    - `models.go` - contains models for mails
    - `router.go` - contains router to handle mail requests
  - `templates` - contains module to handle mail templates
    - `files` - contains mail template files
      - `account-verified.go` - mail template for account verified mail
      - `account-confirmation.go` - mail template for account that needs verification mail
      - `delete-account.go` - mail template for deleted account mail
      - `password-was-changed.go` - mail template for password was changed mail
    - `models.go` - contains models for mail templates
    - `handlers.go` - contains handler to render mail templates
  - `timeouts` - contains module to handle timeout for requests
    - `middleware.go` - contains timeout handler for requests
- `scripts` - contains scripts for build, init and pre-push hook
  - `build.sh` - script to build the project (loop for all files in cmd directory)
  - `init.sh` - script to initialize the project (copy pre-push hook)
  - `postBumpVersion.sh` - script to run after bumping version (set version in docs and push to remote)
  - `postmanSync.sh` - script to sync Postman collection with Postman API
  - `prePush.sh` - script to run before pushing code to remote
  - `bumpVersion.sh` - script to bump version in go.mod file
  - `dropBinaries.sh` - script to remove binaries from bin directory
- `postman` - contains Postman collection for the service
  - `morphyxis-mail-service.postman_collection.json` - Postman collection for the service
- `docker-compose.yml` - docker compose file to run the service in docker (FOR DEVELOPMENT PURPOSE ONLY)
- `Dockerfile` - contains Dockerfiles for building docker images for morpxysis mail service
- `configs` - contains configuration files for air (live reload tool)
  - `air/morphyxis-mail-service.toml` - configuration file for air to run the mail service on air
  - `air/mail-templates-sandbox.toml` - configuration file for air to run the mail templates sandbox on air
- `Makefile` - makefile to run common commands like build, clean, etc.
  - `make mail-service` - to run the mail service on air
  - `make mail-templates-sandbox` - to run the mail templates sandbox on air
  - `make build` - to build all services
  - `make clean` - to clean the project (remove binaries)
  - `make init` - to initialize the project (copy pre-push hook)
  - `make bump-version` - to bump version and push git tag to remote
  - `make help` - to show help message with all available commands
- `go.mod` - go module file
- `go.sum` - go module dependencies file

## GO client - morphyxis-mail-client

client to interact with morphyxis mail service

### Usage

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
