package arraynslice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("test 2 slice return 1 slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 1, 1, 1, 1}
	sumOfSlice := Sum(numbers)
	fmt.Println(sumOfSlice)
	// Output: 5
}
