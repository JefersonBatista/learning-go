# Learning Go

The idea of this repository is just to practice with the Go language and explore its features.

## Number Sequence

Given an initial sequence ($n_0, n_1, n_2, ...$) and multipliers ($x_0, x_1, x_2, ...$), it is a number sequence of the form: \
$n_0, n_1, n_2, ..., n_i = x_0 + x_1*n_{i-1} + x_2*n_{i-2} + ...$

## Kafka

Code using Apache Kafka needs a setup, based on [Kafka Quick Start](https://kafka.apache.org/43/getting-started/quickstart).

```sh
docker run -d -p 9092:9092 --name broker apache/kafka:latest
docker exec --workdir /opt/kafka/ -it broker sh
./bin/kafka-topics.sh --bootstrap-server localhost:9092 \
--create --topic numeric-topic
exit
```
