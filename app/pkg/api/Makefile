update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/gorilla/handlers
	go get -u github.com/gorilla/mux
	go get -u github.com/nndergunov/deliveryApp/app/pkg/logger

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
