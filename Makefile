GOVERSION := $(shell go version | cut -d ' ' -f 3 | cut -d '.' -f 2)

.PHONY: test test-race test-cover-html help
.DEFAULT_GOAL := help

test: ## Run tests
	go test ./cast/...
	go test ./gflag/...
	go test ./peek/...
	go test ./env/...
	go test ./mapping/...
	go test ./path/...
	go test ./properties/...
	go test ./buffer/...
	go test ./hash/...
	go test ./hashmap/...
	go test ./cache/...


test-race: ## Run tests with race detector
	go test -race ./cast/...
	go test -race ./gflag/...
	go test -race ./peek/...
	go test -race ./env/...
	go test -race ./mapping/...
	go test -race ./path/...
	go test -race ./properties/...
	go test -race ./buffer/...
	go test -race ./hash/...
	go test -race ./hashmap/...
	go test -race ./cache/...

test-cover-html: ## Generate test coverage report
	go test ./cast/... -coverprofile=cast_coverage.out -covermode=count
	go test ./gflag/... -coverprofile=gflag_coverage.out -covermode=count
	go test ./peek/... -coverprofile=peek_coverage.out -covermode=count
	go test ./env/... -coverprofile=env_coverage.out -covermode=count
	go test ./mapping/... -coverprofile=mapping_coverage.out -covermode=count
	go test ./path/... -coverprofile=path_coverage.out -covermode=count
	go test ./properties/... -coverprofile=properties_coverage.out -covermode=count
	go test ./buffer/... -coverprofile=buffer_coverage.out -covermode=count
	go test ./hash/... -coverprofile=hash_coverage.out -covermode=count
	go test ./hashmap/... -coverprofile=hashmap_coverage.out -covermode=count
	go test ./cache/... -coverprofile=cache_coverage.out -covermode=count
	go tool cover -func=cast_coverage.out
	go tool cover -func=gflag_coverage.out
	go tool cover -func=peek_coverage.out
	go tool cover -func=env_coverage.out
	go tool cover -func=mapping_coverage.out
	go tool cover -func=path_coverage.out
	go tool cover -func=properties_coverage.out
	go tool cover -func=buffer_coverage.out
	go tool cover -func=hash_coverage.out
	go tool cover -func=hashmap_coverage.out
	go tool cover -func=cache_coverage.out

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
