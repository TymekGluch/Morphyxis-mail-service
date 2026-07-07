-include .env

.PHONY: init
init:
	bash scripts/init.sh

.PHONY: mail-service
mail-service:
	air -c configs/air/morphyxis-mail-service.toml

.PHONY: mail-templates-sandbox
mail-templates-sandbox:
	air -c configs/air/mail-templates-sandbox.toml

.PHONY: docs
docs:
	swag init -g cmd/morphixis-mail-service/main.go -o internal/docs

.PHONY: postman-sync
postman-sync:
	bash scripts/postmanSync.sh

.PHONY: build
build: 
	bash scripts/build.sh

.PHONY: clean
clean-bin:
	bash scripts/dropBinaries.sh

.PHONY: bump-version
bump-version:
	bash scripts/bumpVersion.sh

.PHONY: help
help: 
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build   - Build binaries"
	@echo "  clean-bin - Remove binaries from bin directory"
	@echo "  bump-version - Bump version based on conventional commits"
	@echo "  mail-service - Start the mail service with air"
	@echo "  mail-templates-sandbox - Start the mail templates sandbox with air"
	@echo "  postman-sync - Sync Postman collection"
	@echo "  init - Initialize the project"
	@echo "  docs - Generate Swagger documentation"
	@echo "  help - Show this help message"