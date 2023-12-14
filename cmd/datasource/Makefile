start:
	docker compose up -d --build

down:
	docker compose down

rebuild_datasource_service:
	docker compose up -d --no-deps --build datasource

restart: down start