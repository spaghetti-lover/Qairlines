postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root qairline
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/qairline?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/qairline?sslmode=disable" -verbose down 1
migrateforce:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/qairline?sslmode=disable" -verbose force $(VERSION)
migratedrop:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/qairline?sslmode=disable" -verbose drop
migrategoto:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/qairline?sslmode=disable" -verbose goto $(VERSION)
createmigrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir db/migrations -seq $$name
test:
	go test -v -cover -short ./...
sqlc:
	sqlc generate
server:
	go run cmd/api/main.go
mock:
	mockgen -package mockdb  -destination internal/domain/mock/store.go github.com/spaghetti-lover/qairlines/db/sqlc Store
mail:
	go run cmd/mail/main.go
redis:
	docker run --name redis -p 6379:6379 -d redis:8-alpine