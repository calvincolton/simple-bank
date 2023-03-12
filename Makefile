## postgres: create a new docker database image
postgres/create:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

## postgres: create a new docker database image
postgres/start:
	docker start postgres12

## createdb: create a new database
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

## dropdb: drop existing database
dropdb:
	docker exec -it postgres12 dropdb simple_bank

## migrateup
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

## migratedown
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

## sqlc 
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres/run postgres/start createdb dropdb migrateup migratedown sqlc test
