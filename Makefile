DOCKER_APP_IMG="anti-bruteforce"

build:
	go build -o ./bin/app ./cmd/app

test:
	rm -rf internal/mocks
	docker run --rm -v "$(PWD)":/src -w /src vektra/mockery
	go test -race -count 100 ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.62.0

lint: install-lint-deps
	golangci-lint run ./...

docker-build:
	docker build \
		-t $(DOCKER_APP_IMG) \
		-f build/app/Dockerfile .

docker-run: docker-build
	docker run --rm $(DOCKER_APP_IMG)

up:
	docker compose -f deployments/docker-compose.yml up

down:
	docker compose -f deployments/docker-compose.yml -d down

.PHONY: build