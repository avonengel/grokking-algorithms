package chapter4

func sum(input []int) int {
	if len(input) == 0 {
		return 0
	} else {
		return input[0] + sum(input[1:])
	}
}
