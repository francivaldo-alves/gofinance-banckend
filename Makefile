createDb:
 createdb --username=postgres --owner=postgres go_finance

postgres:
docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

migrateup:
migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up


migratedrop:
migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop

.PRONY: createDb postgres migrateup migratedrop