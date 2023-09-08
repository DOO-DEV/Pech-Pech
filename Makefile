migration-up:
	@echo "Migrating up..."
	go run cmd/migrate/main.go up

migration-down:
	@echo Migration down...
	go run cmd/migrate/main.go down

dev:
	go run cmd/api/main.go