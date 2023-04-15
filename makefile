.PHONY: create_db
create_db:
	@docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres  
.PHONY: migrate_up
migrate_up:
	migrate -path internal/repository/postgre/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up