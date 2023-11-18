run:
	export PORT=8005 && \
	export KRATOS_URL=http://127.0.0.1:4434 && \
	go run .

start:
	docker compose\
		--file docker-compose.yml\
		--env-file=../.env\
		up -d --build

down:
	docker compose\
		--file docker-compose.yml\
		--env-file=../.env\
		down

restart: down start
