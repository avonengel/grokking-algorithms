package chapter4

import "testing"

func TestLenBaseCase(t *testing.T) {
	// given
	var input []int
	wanted := 0

	// when
	got := length(input)

	// then
	if got != wanted {
		t.Errorf("got %d, wanted %d", got, wanted)
	}
}

func TestLen1(t *testing.T) {
	// given
	input := []int{1}
	wanted := 1

	// when
	got := length(input)

	// then
	if got != wanted {
		t.Errorf("got %d, wanted %d", got, wanted)
	}
}
func TestLen2(t *testing.T) {
	// given
	input := []int{1, 2}
	wanted := 2

	// when
	got := length(input)

	// then
	if got != wanted {
		t.Errorf("got %d, wanted %d", got, wanted)
	}
}

func length(input []int) int {
	if len(input) > 0 {
		return 1 + length(input[1:])
	}
	return 0
}
