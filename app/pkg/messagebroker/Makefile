run-mb:
	docker-compose up -d

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/rabbitmq/amqp091-go
	go get -u github.com/nndergunov/deliveryApp/app/pkg/api # needed for tests

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm rabbitmq:3.9.11-management-alpine

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start rabbitmq

docker-stop:
	docker container stop rabbitmq