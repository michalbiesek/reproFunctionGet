.PHONY: start stop restart clean run

start:
	docker compose up -d

stop:
	docker compose down

restart: stop start

clean:
	docker compose down -v

run:
	go run main.go
