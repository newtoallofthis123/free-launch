# Development tasks for free-launch

# Build the binary
build:
	go build -o free-launch .

# Build with optimizations
build-release:
	go build -ldflags="-s -w" -o free-launch .

# Install the binary
install:
	go install .

# Run tests
# Note: This project currently has no tests
# test:
#	go test -v ./...

# Run the application directly
run:
	go run main.go claude

# Run with a specific model
run-model:
	go run main.go claude google/gemma-3

# Clean build artifacts
clean:
	rm -f free-launch

# Format the code
gofmt:
	gofmt -w .

# Lint the code
golint:
	golangci-lint run

# Vet the code
govet:
	go vet ./...

# Check for dependencies
gomod:
	go mod tidy
	go mod verify

# Development workflow (format, build, install)
dev: gofmt build install

# Release workflow (format, lint, vet, build, test)
release: gofmt golint govet gomod build

# Help
default:
	@echo "Available tasks:"
	@echo "  build          - Build the binary"
	@echo "  build-release  - Build with optimizations"
	@echo "  install        - Install the binary"
	@echo "  run            - Run the application directly"
	@echo "  run-model      - Run with a specific model"
	@echo "  clean          - Clean build artifacts"
	@echo "  gofmt          - Format the code"
	@echo "  golint         - Lint the code"
	@echo "  govet          - Vet the code"
	@echo "  gomod          - Check dependencies"
	@echo "  dev            - Development workflow"
	@echo "  release        - Release workflow"
	@echo "  default        - Show this help"