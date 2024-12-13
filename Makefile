.PHONY: build
.default := build

build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose down

rmi_last_image:
	docker rmi $(shell docker images -q | head -n 1)

rm_last_container:
	docker rm -f $(shell docker ps -a -q | head -n 1)