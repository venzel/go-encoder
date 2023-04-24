server:
	go run framework/cmd/server/server.go

test:
	go test -cover ./...

up:
	docker-compose up -d

down:
	docker-compose down

.PHONY: server test up down