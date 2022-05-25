docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start consumer_service_db

docker-stop:
	docker container stop consumer_service_db

