package main

import (
	"fmt"
	"learning-go/src/kafka"
	"learning-go/src/number_sequence"
)

func main() {
	numberSequenceSample()
}

func numberSequenceSample() {
	size := 20
	digits := 5

	printNumber := func(number int) {
		fmt.Printf("%*d ", digits, number)
	}

	printSequence := func(title string, sequence *number_sequence.NumberSequence) {
		fmt.Printf("%s:\n", title)
		for range size {
			printNumber(sequence.Next())
		}
		fmt.Println()
	}

	triangularSequence := number_sequence.Create([]int{1}, 1, 2, -1)
	printSequence("Triangular numbers", triangularSequence)

	fibonacciSequence := number_sequence.Create([]int{1}, 0, 1, 1)
	printSequence("Fibonacci numbers", fibonacciSequence)

	testSequence := number_sequence.Create([]int{4, 2, 3}, -3, 1, -2, 1)
	printSequence("Sequence for test", testSequence)
}

func kafkaSample() {
	go kafka.RunConsumer()
	kafka.RunProducer(number_sequence.Create([]int{1}, 0, 1, -1), 20)
}
