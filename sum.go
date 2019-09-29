package arraynslice

// Sum only return total sum of array value
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll sum the slice and combine into single slice
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, number := range numbersToSum {
		sums[i] = Sum(number)
	}

	return sums
}
