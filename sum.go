package arraynslice

// Sum func receive array and returns sum of arrays members
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll receive any number of slice and return one slice
func SumAll(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
