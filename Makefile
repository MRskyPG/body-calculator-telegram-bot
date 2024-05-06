include .env
.PHONY: migrate
migrate:
	 docker build -t body-calculator-bot-image .
build:
	 go mod download
	 docker run --name body-calculator-database -p 5439:5432 -d --rm -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DATABASE) -e POSTGRES_USER=$(POSTGRES_USER) body-calculator-bot-image
run:
	go run ./cmd/main.go
stop:
	docker stop body-calculator-database
migrate_down:
	docker rmi body-calculator-bot-image
