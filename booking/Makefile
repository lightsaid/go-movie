## postgresql DSN
DATABASE_URL=postgresql://postgres:abc123@localhost:5432/db_movie?sslmode=disable

## migrate: 生成迁移sql文件, exp: make migrate NAME=init_db
migrate:
	migrate create -seq -ext=.sql -dir=./db/migrations $$NAME

## migrate_up: 向上迁移
migrate_up:
	migrate -database ${DATABASE_URL} -path ./db/migrations -verbose up 1

## migrate_down: 向下迁移
migrate_down:
	migrate -database ${DATABASE_URL} -path ./db/migrations -verbose down 1

## 例如： make migrate_force V=1
migrate_force:
	migrate -database ${DATABASE_URL} -path ./db/migrations force $$V

## sqlc generate:
sqlc:
	sqlc generate

## test 测试代码:
test:
	go test -v -cover ./...

api:
	go run ./cmd/api

web:
	go run ./cmd/web

godb:
	docker exec -it postgres-14-alpine psql -U postgres -d db_movie /bin/sh

mock:
	mockgen -package mockdb -destination db/mock/store.go lightsaid.com/db_movie/db/sqlc Store