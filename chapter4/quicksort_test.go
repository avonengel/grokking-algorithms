package chapter4

import (
	"fmt"
	"slices"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{2, 1}},
		{[]int{3, 2, 1}},
		{[]int{15, 23, 4, 29, 10, 18, 30, 56}},
	}

	for _, testData := range tests {
		t.Run(fmt.Sprintf("%v", testData), func(t *testing.T) {
			quicksort(testData.input)
			// create a second slice to compare
			wanted := make([]int, len(testData.input))
			copy(wanted, testData.input)
			slices.Sort(wanted)
			for i, gotI := range testData.input {
				if gotI != wanted[i] {
					t.Errorf("[%d] got: %v, wanted %d", i, gotI, wanted[i])
				}
			}
		})
	}
}

func quicksort(data []int) {
	if len(data) > 2 {
		pivot := data[0]
		// a real implementation would move things around in order to avoid creating new arrays/slices
		less := make([]int, 0, len(data))
		greater := make([]int, 0, len(data))
		for idx, val := range data {
			if idx == 0 {
				continue
			}
			// TODO extend this check to include val == pivot && idx != pivotIdx
			if val < pivot {
				less = append(less, val)
			} else {
				greater = append(greater, val)
			}
		}
		quicksort(less)
		quicksort(greater)
		for i, _ := range less {
			data[i] = less[i]
		}
		data[len(less)] = pivot
		for i := range greater {
			data[len(less)+i] = greater[i]
		}
	} else if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
	}
}
