package quicksort

import "math/rand/v2"

// Use Quicksort algorithm to sort a slice without creating a copy
func Quicksort(slice []int) {
	if len(slice) < 2 {
		return
	}

	pivotIndex := partition(slice)
	Quicksort(slice[:pivotIndex])
	Quicksort(slice[pivotIndex+1:])
}

func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

/* Ensure the following properties of the resultant state of arr:
 * All values stricly smaller than the pivot before it
 * All values equal to or bigger than the pivot after it
 * Return the pivot final index
 */
func partition(slice []int) (pivotIndex int) {
	pivotInitialIndex := rand.IntN(len(slice))
	pivot := slice[pivotInitialIndex]

	numSmallers := 0
	for _, n := range slice {
		if n < pivot {
			numSmallers++
		}
	}

	pivotIndex = numSmallers
	swap(slice, pivotInitialIndex, pivotIndex)

	numSwaps := 0
	for i := range numSmallers {
		if slice[i] >= pivot {
			numSwaps++
		}
	}

	swapCount := 0
	i, j := 0, pivotIndex+1

	for swapCount < numSwaps {
		for slice[i] < pivot {
			i++
		}
		for slice[j] >= pivot {
			j++
		}

		swap(slice, i, j)
		swapCount++
		i++
		j++
	}

	return
}
