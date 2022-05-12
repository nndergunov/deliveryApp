status-check:
	-curl http://localhost:8084/status

cl-start:
	make -C database docker-build
	go run main.go

docker-build:
	make -C database docker-build
	docker-compose up -d

docker-delete:
	make -C database docker-delete
	docker-compose down --remove-orphans --volumes
	docker image rm kitchen_app

docker-rebuild:
	make -C database docker-rebuild
	make docker-delete
	make docker-build

docker-start:
	make -C database docker-start
	docker container start kitchen-service

docker-stop:
	make -C database docker-stop
	docker container stop kitchen-service

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
