up:
	go mod tidy
	go mod vendor
	docker compose up

build:
	docker compose build
