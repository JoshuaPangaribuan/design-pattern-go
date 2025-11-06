.PHONY: help build test fmt clean run-all

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build all patterns
	@echo "Building all patterns..."
	@go build ./...

test: ## Run all tests
	@echo "Running tests..."
	@go test ./...

fmt: ## Format all Go files
	@echo "Formatting code..."
	@go fmt ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@go clean ./...

run-all: ## Run all pattern examples
	@echo "Running all creational patterns..."
	@for dir in creational/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			echo "\n=== Running $$dir ==="; \
			(cd "$$dir" && go run main.go); \
		fi \
	done
	@echo "\nRunning all structural patterns..."
	@for dir in structural/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			echo "\n=== Running $$dir ==="; \
			(cd "$$dir" && go run main.go); \
		fi \
	done
	@echo "\nRunning all behavioral patterns..."
	@for dir in behavioral/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			echo "\n=== Running $$dir ==="; \
			(cd "$$dir" && go run main.go); \
		fi \
	done

