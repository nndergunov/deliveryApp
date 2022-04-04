check-all:
	curl http://localhost:8080/status & \
	curl http://localhost:8081/status & \
	curl http://localhost:8082/status & \
	curl http://localhost:8083/status & \
	curl http://localhost:8084/status & \
	curl http://localhost:8085/status & \
	curl http://localhost:8086/status

docker-build-all:
	cd ./services; \
    	cd ./accounting/ ; make docker-build ; cd ../ ; \
    	cd ./client/ 	 ; make docker-build ; cd ../ ; \
    	cd ./courier/    ; make docker-build ; cd ../ ; \
    	cd ./delivery/   ; make docker-build ; cd ../ ; \
    	cd ./kitchen/ 	 ; make docker-build ; cd ../ ; \
    	cd ./order/ 	 ; make docker-build ; cd ../ ; \
    	cd ./restaurant/ ; make docker-build

docker-delete-all:
	cd ./services ; \
    	cd ./accounting/ ; make docker-delete ; cd ../ ; \
    	cd ./client/ 	 ; make docker-delete ; cd ../ ; \
    	cd ./courier/    ; make docker-delete ; cd ../ ; \
    	cd ./delivery/   ; make docker-delete ; cd ../ ; \
    	cd ./kitchen/ 	 ; make docker-delete ; cd ../ ; \
    	cd ./order/ 	 ; make docker-delete ; cd ../ ; \
    	cd ./restaurant/ ; make docker-delete

docker-rebuild-all:
	cd ./services ; \
    	cd ./accounting/ ; make docker-rebuild ; cd ../ ; \
    	cd ./client/ 	 ; make docker-rebuild ; cd ../ ; \
    	cd ./courier/    ; make docker-rebuild ; cd ../ ; \
    	cd ./delivery/   ; make docker-rebuild ; cd ../ ; \
    	cd ./kitchen/    ; make docker-rebuild ; cd ../ ; \
    	cd ./order/ 	 ; make docker-rebuild ; cd ../ ; \
    	cd ./restaurant/ ; make docker-rebuild

docker-start-all:
	cd ./services ; \
    	cd ./accounting/ ; make docker-start ; cd ../ ; \
    	cd ./client/ 	 ; make docker-start ; cd ../ ; \
    	cd ./courier/ 	 ; make docker-start ; cd ../ ; \
    	cd ./delivery/   ; make docker-start ; cd ../ ; \
    	cd ./kitchen/    ; make docker-start ; cd ../ ; \
    	cd ./order/ 	 ; make docker-start ; cd ../ ; \
    	cd ./restaurant/ ; make docker-start

docker-stop-all:
	cd ./services ; \
    	cd ./accounting/ ; make docker-stop ; cd ../ ; \
    	cd ./client/ 	 ; make docker-stop ; cd ../ ; \
    	cd ./courier/    ; make docker-stop ; cd ../ ; \
    	cd ./delivery/   ; make docker-stop ; cd ../ ; \
    	cd ./kitchen/    ; make docker-stop ; cd ../ ; \
    	cd ./order/ 	 ; make docker-stop ; cd ../ ; \
    	cd ./restaurant/ ; make docker-stop

cl-start-all:
	cd ./services ; \
		cd ./accounting/ ; make cl-start & cd ../ ; \
		cd ./client/ 	 ; make cl-start & cd ../ ; \
		cd ./courier/ 	 ; make cl-start & cd ../ ; \
		cd ./delivery/   ; make cl-start & cd ../ ; \
		cd ./kitchen/	 ; make cl-start & cd ../ ; \
		cd ./order/ 	 ; make cl-start & cd ../ ; \
		cd ./restaurant/ ; make cl-start

update-gomod-all:
	cd ./services ; \
    	cd ./accounting/ ; make update-gomod ; cd ../ ; \
    	cd ./client/	 ; make update-gomod ; cd ../ ; \
    	cd ./courier/ 	 ; make update-gomod ; cd ../ ; \
    	cd ./delivery/   ; make update-gomod ; cd ../ ; \
    	cd ./kitchen/ 	 ; make update-gomod ; cd ../ ; \
    	cd ./order/ 	 ; make update-gomod ; cd ../ ; \
    	cd ./restaurant/ ; make update-gomod

gomod-tidy-all:
	cd ./services ; \
        cd ./accounting/ ; make gomod-tidy ; cd ../ ; \
        cd ./client/	 ; make gomod-tidy ; cd ../ ; \
        cd ./courier/  	 ; make gomod-tidy ; cd ../ ; \
        cd ./delivery/   ; make gomod-tidy ; cd ../ ; \
        cd ./kitchen/	 ; make gomod-tidy ; cd ../ ; \
        cd ./order/		 ; make gomod-tidy ; cd ../ ; \
        cd ./restaurant/ ; make gomod-tidy
