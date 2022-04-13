status-check:
	-curl http://localhost:8081/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm delivery_app

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start delivery-service

docker-stop:
	docker container stop delivery-service

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/nndergunov/deliveryApp/app/pkg/api
	go get -u github.com/nndergunov/deliveryApp/app/pkg/configreader
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger
	go get -u github.com/nndergunov/deliveryApp/app/pkg/server
	go get -u github.com/gorilla/mux

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
