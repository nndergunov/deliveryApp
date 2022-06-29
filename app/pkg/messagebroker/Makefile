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
