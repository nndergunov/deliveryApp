check-all:
	cd ./accounting/ ; make status-check ; cd ../ ; \
    cd ./delivery/   ; make status-check ; cd ../ ; \
    cd ./restaurant/ ; make status-check ; cd ../ ; \
    cd ./order/ 	 ; make status-check

docker-build-all:
	make -C accounting docker-build
	make -C consumer docker-build
	make -C courier docker-build
	make -C delivery docker-build
	make -C kitchen docker-build
	make -C order docker-build
	make -C restaurant docker-build


docker-delete-all:
	make -C accounting docker-delete
	make -C consumer docker-delete
	make -C courier docker-delete
	make -C delivery docker-delete
	make -C kitchen docker-delete
	make -C order docker-delete
	make -C restaurant docker-delete


docker-rebuild-all:
	make -C accounting docker-rebuild
	make -C consumer docker-rebuild
	make -C courier docker-rebuild
	make -C delivery docker-rebuild
	make -C kitchen docker-rebuild
	make -C order docker-rebuild
	make -C restaurant docker-rebuild

docker-start-all:
	make -C accounting docker-start
	make -C consumer docker-start
	make -C courier docker-start
	make -C delivery docker-start
	make -C kitchen docker-start
	make -C order docker-start
	make -C restaurant docker-start

docker-stop-all:
	make -C accounting docker-stop
	make -C consumer docker-stop
	make -C courier docker-stop
	make -C delivery docker-stop
	make -C kitchen docker-stop
	make -C order docker-stop
	make -C restaurant docker-stop

cl-start-all:
	cd ./accounting/ ; make cl-start & cd ../ ; \
	cd ./delivery/   ; make cl-start & cd ../ ; \
	cd ./restaurant/ ; make cl-start & cd ../ ; \
	cd ./order/ 	 ; make cl-start

update-gomod-all:
	cd ./accounting/ ; make update-gomod ; cd ../ ; \
    cd ./delivery/   ; make update-gomod ; cd ../ ; \
    cd ./restaurant/ ; make update-gomod ; cd ../ ; \
    cd ./order/ 	 ; make update-gomod

download-gomod-all:
	cd ./accounting/ ; make download-gomod ; cd ../ ; \
    cd ./delivery/   ; make download-gomod ; cd ../ ; \
    cd ./restaurant/ ; make download-gomod ; cd ../ ; \
    cd ./order/ 	 ; make download-gomod

gomod-tidy-all:
	cd ./accounting/ ; make gomod-tidy ; cd ../ ; \
    cd ./delivery/   ; make gomod-tidy ; cd ../ ; \
    cd ./restaurant/ ; make gomod-tidy ; cd ../ ; \
    cd ./order/		 ; make gomod-tidy
