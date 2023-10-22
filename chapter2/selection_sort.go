package selection_sort

import (
	"math"
	"slices"
)

func selection_sort(slice []int) {
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
	smallestIdx := -1
	smallestValue := math.MaxInt
	for idx, value := range source {
		if value < smallestValue {
			smallestIdx = idx
			smallestValue = value
		}
	}
	return smallestIdx, smallestValue
}
