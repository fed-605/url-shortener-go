run:
	go run ./cmd/url-shortener/main.go

migrate-up:
	@migrate -path migrations -database ${DB_CONN_STR} up

migrate-version:
	@migrate -path migrations -database  ${DB_CONN_STR} version

migrate-down:
	@migrate -path migrations -database ${DB_CONN_STR} down

docker-up:
	docker compose up --build

docker-down:
	docker compose down
