package chapter4

import (
	"errors"
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		input  []int
		wanted int
		err    error
	}{
		{nil, 0, ErrEmpty},
		{[]int{1}, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 5, nil},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v == %d", test.input, test.wanted), func(t *testing.T) {
			got, err := maxRecursive(test.input)
			if got != test.wanted {
				t.Errorf("%v: got %d, wanted %d", test.input, got, test.wanted)
			}
			if !errors.Is(err, test.err) {
				t.Errorf("got err %v, wanted err %v", err, test.err)
			}
		})
	}
}

var ErrEmpty = errors.New("slice is empty")

func maxRecursive(input []int) (int, error) {
	if len(input) > 0 {
		maxTail, err := maxRecursive(input[1:])
		if err != nil {
			return input[0], nil
		}
		if maxTail > input[0] {
			return maxTail, nil
		}
		return input[0], nil
	}
	return 0, ErrEmpty
}
