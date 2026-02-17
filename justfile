# Development tasks for free-launch

# List available tasks
default:
	@just --list

# Build the binary
build:
	go build -o free-launch ./cmd/

# Build with optimizations
build-release:
	go build -ldflags="-s -w" -o free-launch ./cmd/

# Install the binary
install:
	go install ./cmd/

# Run the application directly
run:
	go run ./cmd/ claude

# Run with a specific model
run-model:
	go run ./cmd/ claude google/gemma-3

# Clean build artifacts
clean:
	rm -f free-launch

# Format the code
fmt:
	gofmt -w .

# Lint the code
lint:
	golangci-lint run

# Vet the code
vet:
	go vet ./cmd/...

# Check and tidy dependencies
mod:
	go mod tidy
	go mod verify

# Development workflow (format, build, install)
dev: fmt build install

# Release workflow (format, lint, vet, build)
release: fmt lint vet mod build