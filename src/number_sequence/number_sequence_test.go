package number_sequence

import (
	"testing"
)

func TestNext(t *testing.T) {
	testSeq := Create([]int{4, 2, 3}, -3, 1, -2, 1)
	nextNumbers := []int{4, 2, 3, 0, -7, -7, 4, 8, -10}

	for i, expected := range nextNumbers {
		if number := testSeq.Next(); number != expected {
			t.Errorf("Number %d, index %d in sequence, differs from expected %d.\n", number, i, expected)
		}
	}
}
