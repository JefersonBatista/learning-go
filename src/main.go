package main

import (
	"fmt"
	"learning-go/src/int_sequence"
	"learning-go/src/int_sequence/tail_mult_sequence"
	"learning-go/src/kafka"
	"learning-go/src/quicksort"
	"math/rand/v2"
	"os"
)

func main() {
	options := make(map[string]func())
	options["sequence"] = tailMultSequenceSample
	options["kafka"] = kafkaSample
	options["quicksort"] = quicksortSample

	listOptions := func() {
		for option := range options {
			fmt.Printf("- %s\n", option)
		}
	}

	if len(os.Args) < 2 {
		fmt.Println("Choose what you want to run between. Options:")
		listOptions()
		return
	}

	option := options[os.Args[1]]
	if option != nil {
		option()
	} else {
		fmt.Println("Invalid argument. Use one of:")
		listOptions()
	}
}

func tailMultSequenceSample() {
	size := 20
	digits := 5

	printNumber := func(number int) {
		fmt.Printf("%*d ", digits, number)
	}

	printSequence := func(title string, sequence *tail_mult_sequence.TailMultSequence) {
		fmt.Printf("%s:\n", title)
		for range size {
			printNumber(sequence.Next())
		}
		fmt.Println()
	}

	triangularSequence := tail_mult_sequence.Create([]int{1}, 1, 2, -1)
	printSequence("Triangular numbers", triangularSequence)

	fibonacciSequence := tail_mult_sequence.Create([]int{1}, 0, 1, 1)
	printSequence("Fibonacci numbers", fibonacciSequence)

	testSequence := tail_mult_sequence.Create([]int{4, 2, 3}, -3, 1, -2, 1)
	printSequence("Sequence for test", testSequence)
}

func kafkaSample() {
	go kafka.RunConsumer()
	sequence := tail_mult_sequence.Create([]int{1}, 0, 1, -1)
	condition := func(n int) bool {
		return n >= 0
	}
	nextWithFilter := int_sequence.NextWithFilter(sequence, condition)
	transformer := func(n int) int {
		return 2 * n
	}
	nextAsBytes := func() []byte {
		return int_sequence.ToBytes(transformer(nextWithFilter()))
	}
	kafka.RunProducer(nextAsBytes, 20)
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
