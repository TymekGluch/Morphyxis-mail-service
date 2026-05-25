# MORPHYXIS MAIL SERVICE

Open source go lang package to handle Own Mail Server (Mailcow based) with common mail template and handle common requests like parameter GET request for mails and POST request to send mails

## Get Started

1.  to start working with code, execute command

```bash
make init
```

- to copy scripts to git hooks (pre-push, pre-commit etc.)

## Project Structure

- `github/workflow` - contains github workflow files for CI/CD
- `cmd` - contains files to start the multiple services
  - `morphixis-mail-service.go` - main file to start the mail service
- `internal` - contains internal packages for the service
- `scripts` - contains scripts for build, init and pre-push hook
  - `build.sh` - script to build the project (loop for all files in cmd directory)
  - `init.sh` - script to initialize the project (copy pre-push hook)
  - `prePush.sh` - script to run before pushing code to remote
  - `bumpVersion.sh` - script to bump version in go.mod file
  - `dropBinaries.sh` - script to remove binaries from bin directory
  - `docker-compose.yml` - docker compose file to run the service in docker (FOR DEVELOPMENT PURPOSE ONLY)
  - `configs` - contains configuration files for air (live reload tool)
    - `air/morphyxis-mail-service.toml` - configuration file for air to run the mail service on air
  - `Makefile` - makefile to run common commands like build, clean, etc.
    - `make mail-service` - to run the mail service on air
- `go.mod` - go module file
- `go.sum` - go module dependencies file
