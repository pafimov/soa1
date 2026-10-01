run:
	go run ./cmd/catalog

test:
	go test ./...

build:
	go build ./...

up:
	docker compose up --build

down:
	docker compose down
