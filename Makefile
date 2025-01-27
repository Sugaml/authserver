postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.6

createdb:
	docker exec -it postgres createdb --username=root --owner=root authserver

dropdb:
	docker exec -it postgres dropdb authserver

run:
	go run cmd/main.go

docs:
	swag init -g ./cmd/authserver/main.go -o ./docs --parseInternal true

lint:
	golangci-lint run ./...

mock:
	mockgen -package mockdb -destination internal/core/port/mock/user.go internal/core/port UserRepository

swag:
	swag init -g cmd/main.go -o ./docs --ot go --parseInternal true
