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

go-test-unit:
	go test api/v1/handler/courierhandler/courierhandler_test.go -v
	go test pkg/service/courierservice/courierservice_test.go -v

go-test-integrational:
	go test pkg/storage/courierstorage/courierstorage_integrational_test.go -v

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

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
