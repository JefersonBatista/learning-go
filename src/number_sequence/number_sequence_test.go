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

func TestGetInit(t *testing.T) {
	testSeq := Create([]int{1, 2}, 1, 1)
	gotInit := testSeq.GetInitCopy()

	if len(testSeq.GetInitCopy()) != 2 {
		t.Error("Internal initial size of sequence being modified.")
		return
	}

	gotInit[1] = 1
	if testSeq.GetInitCopy()[1] != 2 {
		t.Error("Internal initial numbers of sequence being modified.")
	}
}

func TestGetTailMult(t *testing.T) {
	testSeq := Create([]int{1}, 0, 1, 2)
	gotTailMult := testSeq.GetTailMultCopy()

	if len(testSeq.GetTailMultCopy()) != 2 {
		t.Error("Internal size of tail multipliers of sequence being modified.")
		return
	}

	gotTailMult[1] = 1
	if testSeq.GetTailMultCopy()[1] != 2 {
		t.Error("Internal tail multipliers of sequence being modified.")
	}
}

func TestGetNextIndex(t *testing.T) {
	testSeq := Create([]int{}, 0)
	for range 4 {
		testSeq.Next()
	}

	if testSeq.GetNextIndex() != 4 {
		t.Error("Wrong next index (outside init).")
	}

	testSeq = Create([]int{1, 1, 1, 1, 1, 1}, 0)
	for range 4 {
		testSeq.Next()
	}

	if testSeq.GetNextIndex() != 4 {
		t.Error("Wrong next index (inside init).")
	}
}
