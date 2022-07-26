push:
	go mod tidy
	golangci-lint run
	gofumpt -l -w .
	git add .
	git commit -m "on main"
	git push origin feature

golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

genservice:
	mockgen -source=pkg/service/accountingservice/accountingservice.go \
    -destination=pkg/mocks/mock_service.go

genstorage:
	mockgen -source=pkg/service/accountingservice/accountingstorage.go \
	-destination=pkg/mocks/mock_storage.go

go-test-unit:
	go test api/v1/handler/accountinghandler/accountinghandler_test.go -v
	go test pkg/service/accountingservice/accountingservice_test.go -v

go-test-integrational:
	go test pkg/storage/accountingstorage/accountingstorage_integrational_test.go -v

go-test-all:
	make go-test-unit
	make go-test-integrational

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
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18

swagger:
	swagger generate spec -o ./api/v1/swagger.yaml --scan-models

grpc:
	 cd api/v1/grpc protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/accounting.proto
