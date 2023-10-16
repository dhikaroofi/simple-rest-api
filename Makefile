.PHONY: build buildAndRunDockerApps migrate-up migrate-down test generate-mock

include .env

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o ./bin/gandiwa

buildAndRunDockerApps:
	@docker compose -f docker-compose.yaml down
	@docker compose -f docker-compose.yaml up --build -d lokadb
	@docker compose -f docker-compose.yaml up --build -d gandiwa

migrate-up:
	@migrate -path ./database/migrations -database $(DATABASE_URL) up

migrate-down:
	@migrate -path ./database/migrations -database $(DATABASE_URL) down

test:
	go fmt ./...
	go test -coverprofile coverage.cov -cover ./...
	go tool cover -func coverage.cov

generate-mock:
	mockery --all


