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

restart: down start


# app_start:
# 	go run ./cmd/analyzer/main.go
