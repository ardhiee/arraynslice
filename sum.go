package arraynslice

// Sum only return total sum of array value
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
