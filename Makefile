migrate_up:
	migrate -source file://internal/infrastructure/database/migrations/ -database "postgresql://postgres:article@localhost:5432/forum_service?sslmode=disable" up

migrate_down:
	migrate -source file://internal/infrastructure/database/migrations/ -database "postgresql://postgres:article@localhost:5432/forum_service?sslmode=disable" down