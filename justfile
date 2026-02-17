# Development tasks for free-launch

# List available tasks
default:
	@just --list

# Build the binary
build:
	go build -o free-launch .

# Build with optimizations
build-release:
	go build -ldflags="-s -w" -o free-launch .

# Install the binary
install:
	go install .

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
fmt:
	gofmt -w .

# Lint the code
lint:
	golangci-lint run

# Vet the code
vet:
	go vet ./...

# Check and tidy dependencies
mod:
	go mod tidy
	go mod verify

# Development workflow (format, build, install)
dev: fmt build install

# Release workflow (format, lint, vet, build)
release: fmt lint vet mod build