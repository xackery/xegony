default: run

all: test vet lint swagger
run:
	go run main.go server

test:
	go test -v -race ./...

fmt:
	gofmt -s -d ./...

lint:
	golint ./...
services: ## Runs docker-compose up, starting database, nats, and web server
	cd ..
	-docker-compose down
	-docker-compose up
	docker-compose down
vet:
	go vet ./...

audit:
	golint ./...
	go vet ./...

cover:
	./misc/scripts/cover.sh
swagger:
	./misc/scripts/swagger.sh

sloc:
	wc -l */**.go

update:
	go get -u ./...

.PHONY: lint test