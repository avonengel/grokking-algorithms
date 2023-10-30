package chapter4

import "testing"

func TestSumBaseCase(t *testing.T) {
	// given
	input := []int{1}
	expected := 1

	// when
	got := sum(input)

	// then
	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}
