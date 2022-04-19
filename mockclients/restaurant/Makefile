cl-start:
	go run main.go

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
