.PHONY:
.run:
	go run cmd/transactional-srv-test/main.go

run: .run

.PHONY: .up
.up:
	docker-compose up -d

up: .up

.PHONY: .migrate
.migrate:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable" goose -dir ./migrations up

migrate: .migrate