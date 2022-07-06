up:
	go mod tidy
	go mod vendor
	docker compose up

build: build-fe build-be

build-be:
	docker compose build

build-fe:
	cd frontend && yarn build

lint-fe:
	cd frontend && yarn lint

fmt-fe:
	cd frontend && yarn fmt

lint-be:
	docker compose run --rm --entrypoint "" app golangci-lint run --enable goimports

fmt-be:
	docker compose run --rm --entrypoint "" app golangci-lint run --enable goimports --fix

lint: lint-fe lint-be

fmt: fmt-fe fmt-be
