LOCAL_BIN:=$(CURDIR)/bin

DEFAULT_GOAL := help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test: ## Test project.
	go test -v ./...

run: ## Run project for local
	go run cmd/bbone/main.go

lint: ## Lint the whole project
	golint ./...

start-db: ## Start by Docker Compose
	docker-compose up

stop-db: ## Stop by Docker Compose
	docker-compose down

# run goose command
# make goose cmd=create
# make goos cmd=up
gc: ## Work with migration
	goose -dir=migrations postgres "postgres://root:root@0.0.0.0:5432/postgres?sslmode=disable" $(cmd)

migrate: ## Migration up
	make gc cmd=up

build: ## Build Project
	go build -ldflags "-s -w" -o bin/bbone cmd/bbone/main.go
