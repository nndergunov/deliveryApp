golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

genstorage:
	mockgen -source=pkg/service/accountingservice/accountingstorage.go \
	-destination=pkg/mocks/mock_storage.go

genservice:
	mockgen -source=pkg/service/accountingservice/accountingservice.go \
    -destination=pkg/mocks/mock_service.go

status-check:
	-curl http://localhost:8080/status

cl-start:
	make -C database docker-build
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm accounting_app

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start accounting-service

docker-stop:
	docker container stop accounting-service

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
