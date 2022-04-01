status-check:
	curl http://localhost:8080/status

run:
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

