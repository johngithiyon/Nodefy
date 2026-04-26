include internal/config/.env
MAIN = cmd/server/main.go

migrate-create:
	migrate create -ext sql -dir internal/repository/storage/migrations/ $(name)

migrate-up:
	migrate -path internal/repository/storage/migrations/ -database "$(CONN_STR)" up

migrate-down:
	migrate -path internal/repository/storage/migrations/ -database "$(CONN_STR)" down 1

migrate-version:
	migrate -path internal/repository/storage/migrations/ -database "$(CONN_STR)" version

migrate-cure:
	migrate -path internal/repository/storage/migrations/ -database "$(CONN_STR)" force $(verno)

start:
	docker start mypostgres
	docker start myredis
	docker start rabbitmq 

run:
	go run $(MAIN)