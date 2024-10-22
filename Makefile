NETWORK_NAME=girlprofile

network:
	@docker network inspect ${NETWORK_NAME} > /dev/null 2>&1 || docker network create ${NETWORK_NAME}

up: network
	@docker compose up -d

down: network
	@docker compose down

reboot: network
	@docker compose down
	@docker compose up -d

logs:
	@docker compose logs -f

ps:
	@docker compose ps
