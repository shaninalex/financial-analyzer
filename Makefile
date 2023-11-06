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

restart: down start
