
# MIGRATIONS
MIGRATIONS_PATH="db/migrations"

# DATABASE
DATABASE_DRIVER ?= postgres
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_USER ?= docker
DATABASE_PASSWORD ?= docker
DATABASE_SSL ?= disable
DATABASE_DATABASE = mvplease
DATABASE_DSN := "${DATABASE_DRIVER}://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DATABASE}?sslmode=${DATABASE_SSL}"

#############
# DATABASE  #
#############
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)

#############
# SERVER    #
#############
.PHONY: run
run:
	go run ./cmd/server/main.go

#############
# TESTS     #
#############
.PHONY: test
test:
	go test ./...

#############
# BUILD     #
#############
.PHONY: build
build:
	go build -o api cmd/server/main.go