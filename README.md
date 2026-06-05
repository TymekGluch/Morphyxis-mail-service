# MORPHYXIS MAIL SERVICE

Small Open source go lang package to handle Own Mail Server (Mailcow based) with common mail template and handle common requests like POST request to send mails with custom templates

## Get Started

1.  to start working with code, execute command

```bash
make init
```

- to copy scripts to git hooks (pre-push, pre-commit etc.)

## GHCR Registry

- `ghcr.io/morphyxis/morphyxis-mail-service` - contains docker images for morphyxis mail service
  - `latest` - latest stable version of the image (tagged with git tag & build manually triggered)
  - `beta` - latest beta version of the image (built on pull request)

to pull the image, execute command

```bash
  docker pull ghcr.io/tymekgluch/morphyxis-mail-service:latest
```

## Project Structure

- `github/workflow` - contains github workflow files for CI/CD
  - `update-image.yml` - pipelines to build and push docker images to ghcr registry
- `cmd` - contains files to start the multiple services
  - `morphixis-mail-service.go` - main file to start the mail service
  - `mail-templates-sandbox` - file to start the mail templates sandbox (FOR DEVELOPMENT PURPOSE ONLY)
- `internal` - contains internal packages for the service
  - `templates` - contains module to handle mail templates
    - `files` - contains mail template files
      - `account-verified.go` - mail template for account verified mail
      - `account-confirmation.go` - mail template for account that needs verification mail
      - `delete-account.go` - mail template for deleted account mail
      - `password-was-changed.go` - mail template for password was changed mail
    - `models.go` - contains models for mail templates
    - `handlers.go` - contains handler to render mail templates
  - `mailcow-integration` - contains integration with Mailcow API (f.e. smtp requests and handler for all mail templates)
    - `client.go` - contains client to interact with mailcow API
    - `config.go` - contains credentials for the mailcow client
    - `constants.go`,
    - `helpers.go` - contains helper functions
    - `models.go` - contains models for mailcow API
  - `router` - contains router for the service
- `scripts` - contains scripts for build, init and pre-push hook
  - `build.sh` - script to build the project (loop for all files in cmd directory)
  - `init.sh` - script to initialize the project (copy pre-push hook)
  - `prePush.sh` - script to run before pushing code to remote
  - `bumpVersion.sh` - script to bump version in go.mod file
  - `dropBinaries.sh` - script to remove binaries from bin directory
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
