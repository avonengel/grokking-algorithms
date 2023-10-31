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
			// create a second slice to compare
			wanted := make([]int, len(testData.input))
			copy(wanted, testData.input)
			slices.Sort(wanted)
			quicksort(testData.input)
			for i, gotI := range testData.input {
				if gotI != wanted[i] {
					t.Errorf("[%d] got: %v, wanted %d", i, gotI, wanted[i])
				}
			}
		})
	}
}

func FuzzQuicksort(f *testing.F) {
	f.Add([]byte{1, 0, 0, 0, 42, 0, 0, 0})
	f.Add([]byte{1, 0, 0, 0, 42, 0, 0})
	f.Fuzz(func(t *testing.T, fuzzInput []byte) {
		toSort := toIntSlice(fuzzInput)
		// create a second slice to compare
		wanted := make([]int32, len(toSort))
		copy(wanted, toSort)
		slices.Sort(wanted)
		quicksort32(toSort)
		for i, gotI := range toSort {
			if gotI != wanted[i] {
				//t.Errorf("[%d] got: %v, wanted %d", i, gotI, wanted[i])
				t.Fatalf("\n  got: %v\nwanted: %v", toSort, wanted)
			}
		}
	})
}

func toIntSlice(input []byte) []int32 {
	result := make([]int32, len(input)/4)
	for resultIndex := 0; resultIndex < len(input)/4; resultIndex++ {
		var value int32 = 0
		for inputIndex := 4 * resultIndex; inputIndex < len(input); inputIndex++ {
			value |= int32(input[inputIndex]) << uint(8*(inputIndex-4*resultIndex))
		}
		result[resultIndex] = value
	}
	return result
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
		for i, v := range less {
			data[i] = v
		}
		data[len(less)] = pivot
		for i, v := range greater {
			data[len(less)+i+1] = v
		}
	} else if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
	}
}
func quicksort32(data []int32) {
	if len(data) > 2 {
		pivot := data[0]
		// a real implementation would move things around in order to avoid creating new arrays/slices
		less := make([]int32, 0, len(data))
		greater := make([]int32, 0, len(data))
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
		quicksort32(less)
		quicksort32(greater)
		for i, v := range less {
			data[i] = v
		}
		data[len(less)] = pivot
		for i, v := range greater {
			data[len(less)+i+1] = v
		}
	} else if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
	}
}
