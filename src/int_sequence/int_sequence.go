package int_sequence

import "fmt"

type IntSequence interface {
	Next() int
}

func ToBytes(number int) []byte {
	return fmt.Append(nil, number)
}

func NextWithFilter(sequence IntSequence, condition func(number int) bool) func() int {
	return func() int {
		next := sequence.Next()
		for ; !condition(next); next = sequence.Next() {
		}
		return next
	}
}
