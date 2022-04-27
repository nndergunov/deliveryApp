status-check:
	-curl http://localhost:8082/status

cl-start:
	make -C database docker-build
	go run main.go

docker-build:
	make -C database docker-build
	docker-compose up -d

docker-delete:
	cd database
	make docker-delete
	cd ../
	docker-compose down --remove-orphans --volumes
	docker image rm restaurant_app

docker-rebuild:
	cd database
	make docker-rebuild
	cd ../
	make docker-delete
	make docker-build

docker-start:
	cd database
	make docker-start
	cd ../
	docker container start restaurant-service

docker-stop:
	cd database
	make docker-stop
	cd ../
	docker container stop restaurant-service

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/nndergunov/deliveryApp/app/pkg/api
	go get -u github.com/nndergunov/deliveryApp/app/pkg/configreader
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger
	go get -u github.com/nndergunov/deliveryApp/app/pkg/server
	go get -u github.com/volatiletech/sqlboiler/v4
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
