status-check:
	curl http://localhost:8086/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm restaurant_app

docker-rebuild:
	make docker-delete 
	make docker-build

docker-start:
	docker container start restaurant-service

docker-stop:
	docker container stop restaurant-service

update-gomod:
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger
	go get -u github.com/nndergunov/deliveryApp/app/pkg/apilib
	make gomod-tidy

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
