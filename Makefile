#
# Export the configuration path as an environment variable
env:
	export CONFIG_PATH="/home/alex/Projects/Golang/tc-chat/config/local.yaml"

# Start the docker-compose services without setting CONFIG_PATH in docker-compose

build:
	docker-compose build
up:
	docker-compose up -d

# Stop the docker-compose services
down:
	docker-compose down

# View the logs of the docker-compose services
logs:
	docker-compose logs -f

# Restart the docker-compose services
restart: down up

.PHONY: env up down logs restart
