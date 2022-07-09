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
	docker compose run \
	--no-deps --rm --entrypoint "" app \
	golangci-lint run --enable goimports --enable lll

fmt-be:
	docker compose run \
	--no-deps --rm --entrypoint "" app \
	bash -c "golines -w -m 80 .; \
	golangci-lint run --enable goimports --fix"

lint: lint-fe lint-be

fmt: fmt-fe fmt-be
