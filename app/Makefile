start-all-cli:
	cd ./services; \
	cd ./accounting/; go run main.go & cd ../; \
	cd ./client/; go run main.go & cd ../; \
	cd ./courier/; go run main.go & cd ../; \
	cd ./delivery/; go run main.go & cd ../; \
	cd ./kitchen/; go run main.go & cd ../; \
	cd ./order/; go run main.go & cd ../; \
	cd ./restaurant/; go run main.go & cd ../;

check-all-local:
	curl http://localhost:8080/status & \
	curl http://localhost:8081/status & \
	curl http://localhost:8082/status & \
	curl http://localhost:8083/status & \
	curl http://localhost:8084/status & \
	curl http://localhost:8085/status & \
	curl http://localhost:8086/status