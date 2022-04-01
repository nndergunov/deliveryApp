status-check:
	curl http://localhost:8083/status

run:
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

