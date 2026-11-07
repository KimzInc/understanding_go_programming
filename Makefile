.DEFAULT_GOAL := build

APP_NAME := main
INPUT_FILE := sample.txt

.PHONY: fmt vet build run clean

# Format Go source files
fmt:
	@echo "Formatting Go files..."
	@go fmt ./...

# Check Go code for potential errors
vet: fmt
	@echo "Vetting Go files..."
	@go vet ./...

# Compile the application
build: vet
	@echo "Building Go application..."
	@go build -o $(APP_NAME) .

# Build and execute the application
run: build
	@echo "Running Go application..."
	@./$(APP_NAME) $(INPUT_FILE)

# Remove compiled application
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)