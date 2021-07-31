.PHONY: run/api
run/api:
	@go run ./cmd/api -db-dsn=${DBDSN}

	current_time = $(shell date +%c)
git_description = $(shell git describe --always --dirty --tags --long)

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags="-s -X 'main.buildTime=${current_time}' -X 'main.version=${git_description}'" -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -X 'main.buildTime=${current_time}' -X 'main.version=${git_description}'" -o=./bin/linux_amd64/api ./cmd/api