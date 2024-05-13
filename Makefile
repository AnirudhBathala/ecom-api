build:
	@go build -o bin/ecom.exe cmd/main.go

run-build: 
	@./bin/ecom.exe

run:
	@go run cmd/main.go

migrate:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrate.go up

migrate-down:
	@go run cmd/migrate/migrate.go down