docker-build-all:
#	docker network create delivery_app_net
#	make -C app/pkg/messagebroker docker-build
#	make -C app/services/accounting docker-build
#	make -C app/services/consumer docker-build
#	make -C app/services/courier docker-build
	make -C app/services/delivery docker-build
#	make -C app/services/kitchen docker-build
#	make -C app/services/order docker-build
#	make -C app/services/restaurant docker-build

docker-delete-all:
#	make -C app/pkg/messagebroker docker-delete
#	make -C app/services/accounting docker-delete
#	make -C app/services/consumer docker-delete
#	make -C app/services/courier docker-delete
	make -C app/services/delivery docker-delete
#	make -C app/services/kitchen docker-delete
#	make -C app/services/order docker-delete
#	make -C app/services/restaurant docker-delete
#	docker network rm delivery_app_net


docker-rebuild-all:
	make -C app/pkg/messagebroker docker-rebuild
	make -C app/services/accounting docker-rebuild
	make -C app/services/consumer docker-rebuild
	make -C app/services/courier docker-rebuild
	make -C app/services/delivery docker-rebuild
	make -C app/services/kitchen docker-rebuild
	make -C app/services/order docker-rebuild
	make -C app/services/restaurant docker-rebuild


docker-start-all:
	make -C app/pkg/messagebroker docker-start
	make -C app/services/accounting docker-start
	make -C app/services/consumer docker-start
	make -C app/services/courier docker-start
	make -C app/services/delivery docker-start
	make -C app/services/kitchen docker-start
	make -C app/services/order docker-start
	make -C app/services/restaurant docker-start


docker-stop-all:
	make -C app/pkg/messagebroker docker-stop
	make -C app/services/accounting docker-stop
	make -C app/services/consumer docker-stop
	make -C app/services/courier docker-stop
	make -C app/services/delivery docker-stop
	make -C app/services/kitchen docker-stop
	make -C app/services/order docker-stop
	make -C app/services/restaurant docker-stop

go-test-unit:
	make -C app/services/accounting go-test-unit
	make -C app/services/consumer go-test-unit
	make -C app/services/courier go-test-unit
	make -C app/services/delivery go-test-unit

go-test-integrational:
	make -C app/services/accounting go-test-integrational
	make -C app/services/consumer go-test-integrational
	make -C app/services/courier go-test-integrational
	make -C app/services/delivery go-test-integrational

go-test-all:
	make go-test-unit
	make go-test-integrational


cl-start-all:
	cd ./app/services/accounting/ ; make cl-start & cd ../ ; \
	cd ./app/services/consumer/ ; make cl-start & cd ../ ; \
	cd ./app/services/courier/ ; make cl-start & cd ../ ; \
	cd ./app/services/delivery/   ; make cl-start & cd ../ ; \
	cd ./app/services/restaurant/ ; make cl-start & cd ../ ; \
	cd ./app/services/order/ 	 ; make cl-start

update-gomod-all:
	cd ./app/services/accounting/ ; make update-gomod ; cd ../ ; \
	cd ./app/services/consumer/ ; make update-gomod ; cd ../ ; \
	cd ./app/services/courier/ ; make update-gomod ; cd ../ ; \
    cd ./app/services/delivery/   ; make update-gomod ; cd ../ ; \
    cd ./app/services/restaurant/ ; make update-gomod ; cd ../ ; \
    cd ./app/services/order/ 	 ; make update-gomod

download-gomod-all:
	cd ./app/services/accounting/ ; make download-gomod ; cd ../ ; \
	cd ./app/services/consumer/ ; make download-gomod ; cd ../ ; \
	cd ./app/services/courier/ ; make download-gomod ; cd ../ ; \
    cd ./app/services/delivery/   ; make download-gomod ; cd ../ ; \
    cd ./app/services/restaurant/ ; make download-gomod ; cd ../ ; \
    cd ./app/services/order/ 	 ; make download-gomod

gomod-tidy-all:
	cd ./app/services/accounting/ ; make gomod-tidy ; cd ../ ; \
	cd ./app/services/consumer/ ; make gomod-tidy ; cd ../ ; \
	cd ./app/services/courier/ ; make gomod-tidy ; cd ../ ; \
    cd ./app/services/delivery/   ; make gomod-tidy ; cd ../ ; \
    cd ./app/services/restaurant/ ; make gomod-tidy ; cd ../ ; \
    cd ./app/services/order/		 ; make gomod-tidy

check-all:
	cd ./app/services/accounting/ ; make status-check ; cd ../ ; \
	cd ./app/services/consumer/ ; make status-check ; cd ../ ; \
	cd ./app/services/courier/ ; make status-check ; cd ../ ; \
    cd ./app/services/delivery/   ; make status-check ; cd ../ ; \
    cd ./app/services/restaurant/ ; make status-check ; cd ../ ; \
    cd ./app/services/order/ 	 ; make status-check
