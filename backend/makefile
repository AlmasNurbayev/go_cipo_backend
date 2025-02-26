# Description: Makefile

cleardb:
	go run cmd/cleardb/main.go -period "1d"

seeds_up:
	go run cmd/seeder/main.go -typeTask "up" -dsn "postgres://postgres:postgres@localhost:5499/cipo_backend?sslmode=disable"

seeds_down:
	go run cmd/seeder/main.go -typeTask "down" -dsn "postgres://postgres:postgres@localhost:5499/cipo_backend?sslmode=disable"

migrate_up:
# TODO - move DSN to env variable or flag
	go run cmd/migrator/main.go -typeTask "up" -dsn "postgres://postgres:postgres@localhost:5499/cipo_backend?sslmode=disable"

migrate_down:
# TODO - move DSN to env variable or flag
	go run cmd/migrator/main.go -typeTask "down" -dsn "postgres://postgres:postgres@localhost:5499/cipo_backend?sslmode=disable"

build:
	go build -o SERVER cmd/server/main.go

run:
	go run cmd/server/main.go

parser:
	go run cmd/parser/main.go

test:
	go test -v -count=1 ./tests/...

docker_run:
	cd ../ && docker-compose up -d --build