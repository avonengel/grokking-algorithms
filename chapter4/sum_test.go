package chapter4

import (
	"fmt"
	"testing"
)

func TestSumBaseCase(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{42}, 42},
	}
	for _, testInput := range tests {
		t.Run(fmt.Sprintf("TestSumBaseCase: %d", testInput.expected), func(t *testing.T) {
			got := sum(testInput.input)
			if got != testInput.expected {
				t.Fatalf("expected %d, got %d", testInput.expected, got)
			}
		})
	}
}
