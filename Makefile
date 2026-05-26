.PHONY: init
init:
	bash scripts/init.sh

.PHONY: mail-service
mail-service:
	air -c configs/air/morphyxis-mail-service.toml

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
	@echo "  help    - Show this help message"