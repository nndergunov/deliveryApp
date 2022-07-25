golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

service:
	mockgen -source=pkg/service/deliveryservice/deliveryservice.go \
    -destination=pkg/mocks/mock_deliveryservice.go

genstorage:
	mockgen -source=pkg/service/deliveryservice/storage.go \
	-destination=pkg/mocks/mock_deliverystorage.go

genclient:
	mockgen -source=pkg/service/clients.go \
	-destination=pkg/mocks/mock_clients.go

go-test-unit:
	go test api/v1/handler/deliveryhandler/deliveryhandler_test.go -v
	go test pkg/service/deliveryservice/deliveryservice_test.go -v

go-test-integrational:
	go test pkg/storage/deliverystorage/deliverystorage_integrational_test.go -v

go-test-all:
	make go-test-unit
	make go-test-integrational

status-check:
	-curl http://localhost:8083/status

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
	go get -u github.com/lib/pq
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18

swagger:
	swagger generate spec -o ./api/v1/swagger.yaml --scan-models
