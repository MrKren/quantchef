up:
	go mod tidy
	go mod vendor
	docker compose up

build: build-fe build-be

build-be:
	docker compose build

build-fe:
	cd frontend && yarn build

watch:
	cd frontend && yarn start

lint-fe:
	cd frontend && yarn lint

fmt-fe:
	cd frontend && yarn fmt

lint-be:
	docker compose run --rm --entrypoint "" app golangci-lint run

lint: lint-fe lint-be

fmt: fmt-fe
