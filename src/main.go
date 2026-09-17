package main

import (
	"fmt"
	"learning-go/src/kafka"
	"learning-go/src/message_sequence"
	"learning-go/src/message_sequence/number_sequence"
	"learning-go/src/quicksort"
	"math/rand/v2"
	"os"
)

func main() {
	options := make(map[string]func())
	options["sequence"] = numberSequenceSample
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

func numberSequenceSample() {
	size := 20
	digits := 5

	printNumber := func(number message_sequence.Message) {
		fmt.Printf("%*d ", digits, number)
	}

	printSequence := func(title string, sequence *number_sequence.Sequence) {
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
	initialSequence := number_sequence.Create([]int{1}, 0, 1, -1)

	condition := func(n message_sequence.Message) bool {
		return n.(number_sequence.Message) >= 0
	}

	transformation := func(n message_sequence.Message) message_sequence.Message {
		return 2 * n.(number_sequence.Message)
	}

	filter := message_sequence.Filter(initialSequence, condition)
	sequence := message_sequence.Transform(filter, transformation)
	kafka.RunProducer(sequence, 20)
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
