server:
	go run framework/cmd/server/server.go

test:
	go test -cover ./...

up:
	docker-compose up -d

down:
	docker-compose down

test:
	cd encoder && go test ./...

.PHONY: server test up down test