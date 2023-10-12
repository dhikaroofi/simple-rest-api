.PHONY: build buildAndRunDockerApps migrate-up migrate-down

MY_ENV_VARIABLE := $(shell echo $$MY_ENV_VARIABLE)
ENV_FILE := .env

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o ./bin/gandiwa

buildAndRunDockerApps:
	@docker compose -f docker-compose.yaml down
	@docker compose -f docker-compose.yaml up --build -d lokadb
	@docker compose -f docker-compose.yaml up --build -d gandiwa

migrate-up:
	@migrate -path ./database/migrations -database "postgres://default:secret@localhost:5432/gandiwa?sslmode=disable" up

migrate-down:
	@migrate -path ./database/migrations -database "postgres://default:secret@localhost:5432/gandiwa?sslmode=disable" down
