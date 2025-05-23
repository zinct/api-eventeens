MYSQL_URL=mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

run:
	go run cmd/app/main.go

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir infrastructure/db/migrations -seq $$name

migrate:
	migrate -path internal/infrastructure/db/migrations -database '$(MYSQL_URL)' up

migrate-down:
	migrate -path internal/infrastructure/db/migrations -database '$(MYSQL_URL)' down 1

migrate-force:
	migrate -path internal/infrastructure/db/migrations -database '$(MYSQL_URL)' force 1
