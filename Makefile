.PHONY: confirm
confirm:
	@echo 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ] 

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

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migrations files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DBDSN} up

.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	migrate -path=./migrations -database=${DBDSN} down

.PHONY: docker/migrations/up
docker/migrations/up:
	docker run -v $(pwd)/migrations:/migrations --network elements_elements-net migrate/migrate -path=/migrations/ -database=${DBDSN} up

.PHONY: docker/migrations/down
docker/migrations/down:
	docker run -v $(pwd)/migrations:/migrations --network elements_elements-net migrate/migrate -path=/migrations/ -database=${DBDSN} down --all

.PHONY: dc/up
dc/up:
	docker-compose up -d

	.PHONY: dc/down
dc/down:
	docker-compose down