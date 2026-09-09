package quicksort

import (
	"math/rand/v2"
	"testing"
)

func TestQuicksort(t *testing.T) {
	size := 100
	slice := make([]int, size)

	for i := range slice {
		slice[i] = rand.IntN(size)
	}

	originalValues := countValues(slice)

	Quicksort(slice)
	for i := range size - 1 {
		if slice[i] > slice[i+1] {
			t.Error("Quicksort did not let the slice ordered.")
			return
		}
	}

	finalValues := countValues(slice)

	for value, count := range originalValues {
		if finalValues[value] != count {
			t.Error("Quicksort did not preserve the values in slice.")
			return
		}
	}
}

func countValues(slice []int) (counting map[int]uint) {
	counting = make(map[int]uint)
	for _, n := range slice {
		counting[n]++
	}
	return
}
