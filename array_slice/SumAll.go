package main

func SumAll(slice1 []int, slice2 []int) []int {
	result := make([]int, 2)

	for _, number := range slice1 {
		result[0] += number
	}

	for _, number := range slice2 {
		result[1] += number
	}

	return result
}
