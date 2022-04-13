check-all:
	-curl http://localhost:8080/status
	-curl http://localhost:8081/status
	-curl http://localhost:8082/status
	-curl http://localhost:8083/status

docker-build-all:
	cd ./accounting/ ; make docker-build ; cd ../ ; \
	cd ./delivery/   ; make docker-build ; cd ../ ; \
	cd ./kitchen/ 	 ; make docker-build ; cd ../ ; \
	cd ./order/ 	 ; make docker-build

docker-delete-all:
	cd ./accounting/ ; make docker-delete ; cd ../ ; \
	cd ./delivery/   ; make docker-delete ; cd ../ ; \
    cd ./kitchen/ 	 ; make docker-delete ; cd ../ ; \
    cd ./order/ 	 ; make docker-delete

docker-rebuild-all:
	cd ./accounting/ ; make docker-rebuild ; cd ../ ; \
    cd ./delivery/   ; make docker-rebuild ; cd ../ ; \
    cd ./kitchen/    ; make docker-rebuild ; cd ../ ; \
    cd ./order/ 	 ; make docker-rebuild

docker-start-all:
	cd ./accounting/ ; make docker-start ; cd ../ ; \
    cd ./delivery/   ; make docker-start ; cd ../ ; \
    cd ./kitchen/    ; make docker-start ; cd ../ ; \
    cd ./order/ 	 ; make docker-start

docker-stop-all:
	cd ./accounting/ ; make docker-stop ; cd ../ ; \
    cd ./delivery/   ; make docker-stop ; cd ../ ; \
    cd ./kitchen/    ; make docker-stop ; cd ../ ; \
    cd ./order/ 	 ; make docker-stop

cl-start-all:
	cd ./accounting/ ; make cl-start & cd ../ ; \
	cd ./delivery/   ; make cl-start & cd ../ ; \
	cd ./kitchen/	 ; make cl-start & cd ../ ; \
	cd ./order/ 	 ; make cl-start

update-gomod-all:
	cd ./accounting/ ; make update-gomod ; cd ../ ; \
    cd ./delivery/   ; make update-gomod ; cd ../ ; \
    cd ./kitchen/ 	 ; make update-gomod ; cd ../ ; \
    cd ./order/ 	 ; make update-gomod

download-gomod-all:
	cd ./accounting/ ; make download-gomod ; cd ../ ; \
    cd ./delivery/   ; make download-gomod ; cd ../ ; \
    cd ./kitchen/ 	 ; make download-gomod ; cd ../ ; \
    cd ./order/ 	 ; make download-gomod

gomod-tidy-all:
	cd ./accounting/ ; make gomod-tidy ; cd ../ ; \
    cd ./delivery/   ; make gomod-tidy ; cd ../ ; \
    cd ./kitchen/	 ; make gomod-tidy ; cd ../ ; \
    cd ./order/		 ; make gomod-tidy
