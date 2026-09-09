compile:
	go build -o bin/main src/main.go

clear:
	rm bin/main

deps:
	go mod tidy

run-code:
	go run ./src $(ARGS)

run-bin:
	./bin/main $(ARGS)

test:
	go test \
		./src/number_sequence \
		./src/quicksort

KAFKA_IMAGE ?= apache/kafka:latest

kafka-up:
	docker run -d -p 9092:9092 --name broker $(KAFKA_IMAGE)
	docker exec --workdir /opt/kafka/ -it broker ./bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--create --topic numeric-topic

kafka-down:
	docker stop broker
	docker rm broker
