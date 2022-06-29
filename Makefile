docker-build-all:
	make -C app/pkg/messagebroker run-mb
	# make -C app/services/accounting docker-build
	# make -C app/services/consumer docker-build
	# make -C app/services/courier docker-build
	# make -C app/services/delivery docker-build
	make -C app/services/kitchen docker-build
	make -C app/services/order docker-build
	make -C app/services/restaurant docker-build
