docker-build-app:
	docker network create delivery_app_net
	make -C app/pkg/messagebroker docker-build
	make -C app/services docker-build-all


docker-delete-all:
	make -C app/pkg/messagebroker docker-delete
	make -C app/services docker-delete-all
	docker network rm delivery_app_net


docker-rebuild-all:
	make -C app/pkg/messagebroker docker-rebuild
	make -C app/services docker-rebuild-all


docker-start-all:
	make -C app/pkg/messagebroker docker-start
	make -C app/services docker-start-all


docker-stop-all:
	make -C app/pkg/messagebroker docker-stop
	make -C app/services docker-stop-all