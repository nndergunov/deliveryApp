status-check:
	curl http://localhost:8082/status

run:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm courier_app

docker-rebuild:
	make docker-delete 
	make docker-build

docker-start:
	docker container start courier-service

docker-stop:
	docker container stop courier-service

