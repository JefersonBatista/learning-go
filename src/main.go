package main

import (
	"fmt"
	"learning-go/src/kafka"
	"learning-go/src/number_sequence"
	"learning-go/src/quicksort"
	"math/rand/v2"
	"os"
)

func main() {
	options := make(map[string]func())
	options["number_sequence"] = numberSequenceSample
	options["kafka"] = kafkaSample
	options["quicksort"] = quicksortSample

	if len(os.Args) < 2 {
		fmt.Println("Choose what you want to run between. Options:")
		for option := range options {
			fmt.Printf("- %s\n", option)
		}
		return
	}

	option := options[os.Args[1]]
	if option != nil {
		option()
	} else {
		fmt.Println("Invalid argument. Use one of:")
		for option := range options {
			fmt.Printf("- %s\n", option)
		}
	}
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
	sequence := number_sequence.Create([]int{1}, 0, 1, -1)
	kafka.RunProducer(sequence.NextAsBytes, 20)
}

func quicksortSample() {
	size := 20
	slice := make([]int, size)

	for i := range size {
		slice[i] = rand.IntN(size)
	}

	fmt.Print("Original slice: ")
	for _, n := range slice {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	quicksort.Quicksort(slice)

	fmt.Print("Ordered slice:  ")
	for _, n := range slice {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
