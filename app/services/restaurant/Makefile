swagger-generate:
	swagger generate spec --scan-models --output=./api/v1/swagger.yaml

golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

status-check:
	-curl http://localhost:8086/status

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
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/adrianbrad/psqldocker
	go get -u github.com/adrianbrad/psqltest
	go get -u github.com/friendsofgo/errors
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq
	go get -u github.com/stretchr/testify
	go get -u github.com/volatiletech/sqlboiler/v4
	go get -u github.com/volatiletech/strmangle

	go get -u github.com/nndergunov/deliveryApp/app/pkg/api
	go get -u github.com/nndergunov/deliveryApp/app/pkg/configreader
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
