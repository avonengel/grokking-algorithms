package selection_sort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// given
	input := testInput()

	// when
	selection_sort(input)

	// then
	wanted := testInput()
	slices.Sort(wanted)
	if !slices.Equal(input, wanted) {
		t.Errorf("\n   got: %v\nwanted: %v", input, wanted)
	}
}

func testInput() []int {
	return []int{6, 9, 4, 23, 11, 8, 2}
}
