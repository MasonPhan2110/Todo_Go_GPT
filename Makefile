postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine3.18
createdb:
	docker exec -it postgres createdb --username=root --owner=root todo_gpt
dropdb:
	docker exec -it postgres dropdb todo_gpt
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_gpt?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_gpt?sslmode=disable" -verbose down
sqlc:
	sqlc generate
sqlc_window: 
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlc_window