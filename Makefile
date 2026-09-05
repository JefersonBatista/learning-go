compile:
	go build -o bin/main src/main.go

clear:
	rm bin/main

deps:
	go mod tidy

run-code:
	go run ./src

run-bin:
	./bin/main

test:
	go test ./src/number_sequence
