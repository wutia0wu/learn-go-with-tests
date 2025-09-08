package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	result := make([]int, 0, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {

	// lengthOfNumbers := len(numbersToSum)
	// result := make([]int, 0, lengthOfNumbers)

	var result []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			// 注意 empty slice 将会导致越界引发 panic
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}

	return result
}
