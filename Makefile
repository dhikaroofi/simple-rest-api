.PHONY: build buildAndRunDockerApps migrate-up migrate-down

include .env

DB_URL :="postgres://default:secret@localhost:5432/gandiwa?sslmode=disable"

# this command build is only can be run from linux
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


