status-check:
	curl http://localhost:8081/status

run:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm client_app

docker-rebuild:
	make docker-delete 
	make docker-build

docker-start:
	docker container start client-service

docker-stop:
	docker container stop client-service

update-gomod:
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger
	go mod tidy -go=1.18 -compat=1.18
