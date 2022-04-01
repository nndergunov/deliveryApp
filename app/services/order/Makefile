status-check:
	curl http://localhost:8085/status

run:
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

