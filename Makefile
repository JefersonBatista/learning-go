compile:
	go build -o bin/main src/main.go

clear:
	rm bin/main

deps:
	go mod tidy
