.PHONY: tidy test

tidy:
	@echo "Running Go mod tidy in backend..."
	go mod tidy

test: tidy
	@echo "Running Go tests in backend..."
	go test -v ./...