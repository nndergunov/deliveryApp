update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u google.golang.org/grpc

	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18

