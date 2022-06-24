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