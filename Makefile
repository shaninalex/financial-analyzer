start:
	docker compose\
		--file docker-compose.yml\
		--env-file=.env\
		up -d --build

down:
	docker compose\
		--file docker-compose.yml\
		--env-file=.env\
		down

down_volumes:
	docker compose\
		--file docker-compose.yml\
		--env-file=.env\
		down -v

rebuild_router:
	docker compose up -d --no-deps --build router

rebuild_datasource:
	docker compose up -d --no-deps --build datasource

rebuild_auth:
	docker compose \
		--file docker-compose.yml\
		--env-file=.env\
		up -d --no-deps --build auth

restart: down start
