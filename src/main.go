package main

import (
	"fmt"
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

	printInit := func(init []int) {
		for _, number := range init {
			printNumber(number)
		}
	}

	printSequence := func(title string, init []int, sequence *number_sequence.NumberSequence) {
		fmt.Printf("%s:\n", title)
		printInit(init)
		for range size - len(init) {
			printNumber(sequence.Next())
		}
		fmt.Println()
	}

	init := []int{1}
	triangularSequence := number_sequence.Create(init, 1, 2, -1)
	printSequence("Triangular numbers", init, triangularSequence)

	init = []int{1}
	fibonacciSequence := number_sequence.Create(init, 0, 1, 1)
	printSequence("Fibonacci numbers", init, fibonacciSequence)

	init = []int{4, 2, 3}
	testSequence := number_sequence.Create(init, -3, 1, -2, 1)
	printSequence("Sequence for test", init, testSequence)
}
