package number_sequence

/* With numbers (n[0], n[1], n[2], ...) as init, and (x[0], x[1], x[2], ...) as multipliers
 * Implement number sequences of the form:
 * n[0], n[1], n[2], ..., n[i] = x[0] + x[1]*n[i-1] + x[2]*n[i-2] + ...
 */

type NumberSequence struct {
	tailMult   []int
	unitMult   int
	tail       []int
	readOffset uint
}

func Create(init []int, multipliers ...int) (newSequence *NumberSequence) {
	initSize := len(init)
	tailSize := len(multipliers) - 1

	newSequence = &NumberSequence{
		tailMult:   multipliers[1:],
		unitMult:   multipliers[0],
		tail:       make([]int, tailSize),
		readOffset: 0,
	}

	initSteps := min(initSize, tailSize)
	initIndex := max(initSize-tailSize, 0)
	tailIndex := tailSize - initSteps

	for range initSteps {
		newSequence.tail[tailIndex] = init[initIndex]
		initIndex++
		tailIndex++
	}

	return
}

// Read index for the way last numbers (the tail) are stored
func (sequence *NumberSequence) calcTailIndex(index int) (tailIndex int) {
	tailSize := len(sequence.tail)

	// To not have to move last numbers to the left, I use an offset read strategy
	adjust := func(index int) (adjustedIndex int) {
		adjustedIndex = (index + int(sequence.readOffset)) % tailSize
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

func (sequence *NumberSequence) Next() (number int) {
	number = sequence.unitMult
	for index, mult := range sequence.tailMult {
		tailIndex := sequence.calcTailIndex(index)
		number += mult * sequence.tail[tailIndex]
	}

	sequence.tail[sequence.readOffset] = number
	sequence.readOffset = (sequence.readOffset + 1) % uint(len(sequence.tail))

	return
}
