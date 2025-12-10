BIN := "./bin"

.PHONY:test
test:
	go test -count 1 -v $(shell go list ./... | grep -v /integration_tests)

test-race:
	go test -race -count 100 -v $(shell go list ./... | grep -v /integration_tests)

lint: 
	golangci-lint run

vet:
	staticcheck ./...

check: lint vet test

run: 
	go run ./cmd/server/

build_ab:
	go build -v -o "$(BIN)/anti_bruteforce" ./cmd/server/

up:
	docker compose up anti_bruteforce -d --build

down:
	docker compose down

# integration-tests:
# 	docker compose up integration-tests --build
