all: unit build run test_newman down

unit:
	go test ./... -coverprofile=coverage.out

build:
	docker compose build

run:
	docker compose up -d

down:
	docker compose down

test_newman:
	cd atdd && newman run healthcheck.success_collection.json -e healthcheck-env.postman_environment.json -d healthcheck.data_success.json
	cd atdd && newman run healthcheck.unsuccess_collection.json -e healthcheck-env.postman_environment.json -d healthcheck.data_unsuccess.json