package binsort

import (
	"errors"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// given
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	// when
	index, error := binarySearch(input, 13)

	// then
	if error != nil {
		t.Fatal(error)
	}
	if index != 12 {
		t.Errorf("got %d, wanted 12", index)
	}
}

func binarySearch(input []int, item int) (int, error) {
	return 0, errors.ErrUnsupported
}
