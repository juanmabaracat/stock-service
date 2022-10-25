run: ## Run the application
	go run ./cmd/main.go

test: ## Run unit tests
	go test `go list ./... | grep -v 'docs'` -race