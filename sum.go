package arraynslice

// Sum func receive array and returns sum of arrays members
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
