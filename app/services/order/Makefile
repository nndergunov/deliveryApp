status-check:
	curl http://localhost:8085/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm order_app

docker-rebuild:
	make docker-delete 
	make docker-build

docker-start:
	docker container start order-service

docker-stop:
	docker container stop order-service

update-gomod:
	go get -u github.com/nndergunov/deliveryApp/app/pkg/api
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger	
	go get -u github.com/nndergunov/deliveryApp/app/pkg/server
	make gomod-tidy

download-gomod:
	go get github.com/nndergunov/deliveryApp/app/pkg/api
	go get github.com/nndergunov/deliveryApp/app/pkg/logger
	go get github.com/nndergunov/deliveryApp/app/pkg/server

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
