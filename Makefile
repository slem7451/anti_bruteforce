DOCKER_APP_IMG="anti-bruteforce"
DOCKER_CLI_IMG="anti-bruteforce-cli"
DOCKER_TEST_IMG="integration-tests"

build:
	go build -o ./bin/app ./cmd/app

test:
	rm -rf internal/mocks
	docker run --rm -v "$(PWD)":/src -w /src vektra/mockery
	go test -race -count 100 ./internal/...

integration-tests:
	docker compose -f deployments/docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from integration-tests
	docker compose -f deployments/docker-compose.test.yml down	

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.62.0

lint: install-lint-deps
	golangci-lint run ./...

docker-build:
	docker build \
		-t $(DOCKER_APP_IMG) \
		-f build/app/Dockerfile .

	docker build \
		-t $(DOCKER_TEST_IMG) \
		-f build/test/Dockerfile .

docker-run: docker-build
	docker run --rm $(DOCKER_APP_IMG)

up:
	docker compose -f deployments/docker-compose.yml up

down:
	docker compose -f deployments/docker-compose.yml -d down

cli-build:
	docker build \
		-t $(DOCKER_CLI_IMG) \
		-f build/cli/Dockerfile .

	docker run --name cli-container -d $(DOCKER_CLI_IMG)
	docker cp cli-container:ab ./bin/ab
	docker stop cli-container && docker rm cli-container

.PHONY: build