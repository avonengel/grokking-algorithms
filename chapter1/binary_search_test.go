package binsearch

import (
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

func BenchmarkWithHighLowComparison(b *testing.B) {
	b.StopTimer()
	input := testInput(0, b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		result, err := binarySearch(input, i)
		if err != nil {
			b.Error(err)
		}
		if result != i {
			b.Errorf("wanted %d, but got %d", i, result)
		}
	}
}

func BenchmarkWithMidComparison(b *testing.B) {
	b.StopTimer()
	input := testInput(0, b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		result, err := binarySearch(input, i)
		if err != nil {
			b.Error(err)
		}
		if result != i {
			b.Errorf("wanted %d, but got %d", i, result)
		}
	}
}
