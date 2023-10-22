package selection_sort

import (
	"slices"
)

func selectionSort(slice []int) {
	// create a second array to select into
	selectFrom := make([]int, len(slice))
	copy(selectFrom, slice)

	// for each index in the second array, select the smallest item from the input
	for i, _ := range slice {
		idx, value := findSmallest(selectFrom)
		slice[i] = value
		selectFrom = slices.Delete(selectFrom, idx, idx+1)
	}
}

func findSmallest(source []int) (int, int) {
	smallestIdx := 0
	smallestValue := source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if value < smallestValue {
			smallestIdx = i
			smallestValue = value
		}
	}
	return smallestIdx, smallestValue
}
