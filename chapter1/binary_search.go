package binsearch

import (
	"errors"
)

func binarySearch(input []int, item int) (int, error) {
	return binarySearchWithMidComparison(input, item)
}
func binarySearchWithHighLowComparison(input []int, item int) (int, error) {
	low := 0
	high := len(input) - 1
	for {
		// DONE: this will repeatedly check the same element, as only low OR high are updated each iteration
		if input[low] == item {
			return low, nil
		} else if input[high] == item {
			return high, nil
		}
		// DONE: this should go into for low < high or similar, but I can't make it work right now - too tired
		if low+1 == high {
			return 0, errors.New("NotFound")
		}
		mid := (high + low) / 2
		if input[mid] <= item {
			low = mid
		} else {
			high = mid
		}
	}
}

func binarySearchWithMidComparison(input []int, item int) (int, error) {
	low := 0
	high := len(input) - 1
	//fmt.Printf("searching for %d\n", item)
	for low <= high {
		mid := (high + low) / 2
		if input[mid] == item {
			return mid, nil
		}
		//fmt.Printf("%d .. %d (%d) .. %d", low, mid, input[mid], high)
		if input[mid] <= item {
			//fmt.Printf(" => low = %d\n", mid+1)
			low = mid + 1
		} else {
			//fmt.Printf(" => high = %d\n", mid-1)
			high = mid - 1
		}
	}
	return 0, errors.New("NotFound")
}
