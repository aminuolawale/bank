postgres:
	docker run --name postgres12 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=${pass} -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

connectdb:
	docker exec -it postgres12 psql simple_bank -U root 
migrateUp:
	migrate -path db/migrations -database "postgresql://root:${pass}@localhost:5431/simple_bank?sslmode=disable" -verbose up
migrateDown:
	migrate -path db/migrations -database "postgresql://root:${pass}@localhost:5431/simple_bank?sslmode=disable" -verbose down
sqlc: 
	sqlc generate


test:
	go test -v -cover ./...