package algo

import (
	"fmt"
	"testing"
	"reflect"
)

func TestSelectionSort(t *testing.T) {
	test := func(input []int, expected []int) {
		fmt.Println("Sorting ", input)
		sorted := SelectionSort(input)
		if !reflect.DeepEqual(sorted, expected) {
			t.Error("Expected ", expected, ", got ", sorted)
		}
	}

	test([]int{}, []int{})
	test([]int{1}, []int{1})
	test([]int{2, 1}, []int{1, 2})
	test([]int{3, 2, 1}, []int{1, 2, 3})
	test([]int{4, 5, 3, 2, 1}, []int{1, 2, 3, 4, 5})
	test([]int{5, 3, 6, 2, 10}, []int{2, 3, 5, 6, 10})
}
