APP_NAME=ecommerce-api
DB_CONTAINER=ecommerce-db-1

up:
	docker compose up --build

dev:
	docker compose up

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker ps

migrate:
	docker exec -it $(DB_CONTAINER) psql -U postgres -d ecommerce -f /db/migrations.sql

db:
	docker exec -it $(DB_CONTAINER) psql -U postgres -d ecommerce

restart:
	docker compose down
	docker compose up --build
