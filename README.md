# Learning Go

The idea of this repository is just to practice with the Go language and explore its features.

## Number Sequence

Given an initial sequence ($n_0, n_1, n_2, ...$) and multipliers ($x_0, x_1, x_2, ...$), it is a number sequence of the form: \
$n_0, n_1, n_2, ..., n_i = x_0 + x_1 * n_{i-1} + x_2 * n_{i-2} + ...$

Usage: `make run-code` (or `run-bin` to use compiled binary) `ARGS=number_sequence`.

## Kafka

Code using Apache Kafka needs a setup (`make kafka-up`, optionally passing a `KAFKA_IMAGE`), based on [Kafka Quick Start](https://kafka.apache.org/43/getting-started/quickstart). \
To unmount kakfa setup: `make kafka-down`.

Usage: `make run-code` (or `run-bin` to use compiled binary) `ARGS=kafka`.
