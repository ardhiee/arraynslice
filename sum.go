package arraynslice

// Sum func receive array and returns sum of arrays members
func Sum(numbers [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}
