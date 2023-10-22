package binsort

import (
	"errors"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// given
	input := testInput(1, 16)

	// when
	index, err := binarySearch(input, 13)

	// then
	if err != nil {
		t.Fatal(err)
	}
	if index != 12 {
		t.Errorf("got %d, wanted 12", index)
	}
}

func testInput(lower int, upper int) []int {
	result := make([]int, upper-lower+1)
	for i := 0; i <= upper-lower; i++ {
		result[i] = lower + i
	}
	return result
}

func TestBinarySearchSingleItemArray(t *testing.T) {
	// given
	input := []int{5}

	// when
	index, err := binarySearch(input, 5)

	// then
	if err != nil {
		t.Fatal(err)
	}
	if index != 0 {
		t.Errorf("got %d, wanted 0", index)
	}
}

func TestNotInArray(t *testing.T) {
	// given
	input := testInput(1, 16)

	// when
	index, err := binarySearch(input, 42)
	// then
	if err == nil || err.Error() != "NotFound" {
		t.Fatalf("Error NotFound expected, got (%d, %s)", index, err)
	}
}
func binarySearch(input []int, item int) (int, error) {
	low := 0
	high := len(input) - 1
	for {
		if input[low] == item {
			return low, nil
		} else if input[high] == item {
			return high, nil
		}
		if low+1 == high {
			return 0, errors.New("NotFound")
		}
		mid := (high-low)/2 + low
		if input[mid] <= item {
			low = mid
		} else {
			high = mid
		}
	}
}
