package number_sequence

import (
	"fmt"
	"learning-go/src/message_sequence"
)

type Message int

func (msg Message) Bytes() []byte {
	return fmt.Append(nil, msg)
}

func (msg Message) String() string {
	return fmt.Sprintf("%d", msg)
}

/* With numbers (n[0], n[1], n[2], ...) as init, and (x[0], x[1], x[2], ...) as multipliers
 * Implement number sequences of the form:
 * n[0], n[1], n[2], ..., n[i] = x[0] + x[1]*n[i-1] + x[2]*n[i-2] + ...
 */
type Sequence struct {
	init       []int
	unitMult   int
	tailMult   []int
	tail       []int
	readOffset int
	nextIndex  int
}

func Create(init []int, unitMult int, tailMult ...int) (newSequence *Sequence) {
	newSequence = &Sequence{
		init:       init,
		tailMult:   tailMult,
		unitMult:   unitMult,
		tail:       make([]int, len(tailMult)),
		readOffset: 0,
		nextIndex:  0,
	}

	if len(tailMult) == 0 {
		newSequence.tailMult = []int{0}
		newSequence.tail = []int{0}
	}

	initSteps := min(len(init), len(tailMult))
	initIndex := max(len(init)-len(tailMult), 0)
	tailIndex := len(tailMult) - initSteps

	for range initSteps {
		newSequence.tail[tailIndex] = init[initIndex]
		initIndex++
		tailIndex++
	}

	return
}

// Read index for the way last numbers (the tail) are stored
func (sequence *Sequence) calcTailIndex(index int) (tailIndex int) {
	tailSize := len(sequence.tail)

	// To not have to move last numbers to the left, I use an offset read strategy
	adjust := func(index int) (adjustedIndex int) {
		adjustedIndex = (index + sequence.readOffset) % tailSize
		return
	}

	/* The first multipliers multiplie the last numbers in sequence,
	 * so the tail index must be reversed
	 */
	reverse := func(index int) (reversedIndex int) {
		reversedIndex = (tailSize - 1) - index
		return
	}

	tailIndex = adjust(reverse(index))
	return
}

func (sequence *Sequence) Next() (message message_sequence.Message) {
	defer func() {
		sequence.nextIndex++
	}()

	if sequence.nextIndex < len(sequence.init) {
		message = Message(sequence.init[sequence.nextIndex])
		return
	}

	number := sequence.unitMult
	for index, mult := range sequence.tailMult {
		tailIndex := sequence.calcTailIndex(index)
		number += mult * sequence.tail[tailIndex]
	}
	message = Message(number)

	sequence.tail[sequence.readOffset] = number
	sequence.readOffset = (sequence.readOffset + 1) % len(sequence.tail)

	return
}

func (sequence *Sequence) GetInitCopy() (initCopy []int) {
	initCopy = make([]int, len(sequence.init))
	copy(initCopy, sequence.init)
	return
}

func (sequence *Sequence) GetUnitMult() int {
	return sequence.unitMult
}

func (sequence *Sequence) GetTailMultCopy() (tailMultCopy []int) {
	tailMultCopy = make([]int, len(sequence.tailMult))
	copy(tailMultCopy, sequence.tailMult)
	return
}

func (sequence *Sequence) GetNextIndex() int {
	return sequence.nextIndex
}
