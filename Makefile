include .env
export

export PROJECT_ROOT=$(shell pwd)

start-postgres:
	@docker compose up -d navigation-postgres-db

finish-postgres:
	@docker compose down navigation-postgres-db

start-redis:
	@docker compose up -d navigation-redis

finish-redis:
	@docker compose down navigation-redis

start-kafka:
	@docker compose up -d navigation-kafka

finish-kafka:
	@docker compose down navigation-kafka

cleanup-postgres:
	@read -p "Очистить pg_data? ВНИМАНИЕ! Опасность утери данных! [y/N]: " choice; \
	if [ "$$choice" = "y" ] || [ "$$choice" = "Y" ]; then \
		docker compose down -v navigation-postgres-db && \
		sudo rm -rf ${PROJECT_ROOT}/out/pg_data && \
		echo "Очищено"; \
	else echo "Операция отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Вы не передали название миграции!" \
		exit 1; \
	fi;
	@docker compose run --rm navigation-postgres-db-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Нет параметра action!"; \
		exit 1; \
	fi;
	@docker compose run --rm navigation-postgres-db-migrate \
	-path /migrations \
	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@navigation-postgres-db:5432/${POSTGRES_DB}?sslmode=disable \
	"$(action)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down


start-log-service:
	@docker compose up -d log --build

finish-log-service:
	@docker compose down log