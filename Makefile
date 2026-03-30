BOT_BINARY=discord-bot
BROKER_BINARY=profile-broker
AUTH_BINARY=auth-service

## Build the bot for running locally on Mac`
build_bot_mac:
	@echo "Building bot for Mac..."
	cd ./discord-bot && go build -o ./dist/$(BOT_BINARY)-mac ./cmd/api
	echo "Done!"

build_broker_mac:
	@echo "Building broker for Mac..."
	cd ./profile-broker && go build -o ./dist/$(BROKER_BINARY)-mac ./cmd/api
	echo "Done!"

build_auth_mac:
	@echo "Building auth for Mac..."
	cd ./auth-service && go build -o ./dist/$(AUTH_BINARY)-mac ./cmd/api
	echo "Done!"

## Build the bot for running in Docker (Linux)
build_bot:
	@echo "Building bot for Docker..."
	cd ./discord-bot && env GOOS=linux CGO_ENABLED=0 go build -o ./dist/$(BOT_BINARY) ./cmd/api
	@echo "Done!"

build_broker:
	@echo "Building broker for Docker..."
	cd ./profile-broker && env GOOS=linux CGO_ENABLED=0 go build -o ./dist/$(BROKER_BINARY) ./cmd/api
	@echo "Done!"

build_auth:
	@echo "Building auth for Docker..."
	cd ./auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ./dist/$(AUTH_BINARY) ./cmd/api
	@echo "Done!"

up_build: build_bot build_broker build_auth
	@echo "Stopping and removing existing containers..."
	docker compose down
	@echo "Starting new containers with the latest builds..."
	docker compose up -d --build
	@echo "Done!"

start:
	@echo "Starting front end app"
	cd ./front-end && yarn && yarn dev

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"