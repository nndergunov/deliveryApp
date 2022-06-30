golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

status-check:
	-curl http://localhost:8082/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm courier_app

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start courier-service

docker-stop:
	docker container stop courier-service

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/nndergunov/deliveryApp/app/pkg/api
	go get -u github.com/nndergunov/deliveryApp/app/pkg/configreader
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger
	go get -u github.com/nndergunov/deliveryApp/app/pkg/server
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
