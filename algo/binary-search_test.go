package algo

import "fmt"
import "testing"

func TestBinarySearch(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println("Searching in ", s)

	searchOne, _ := BinarySearch(s, 1)
	if !searchOne {
		t.Error("Expected found 1 true , got ", searchOne)
	}

	searchFour, _ := BinarySearch(s, 4)
	if searchFour {
		t.Error("Expected found 4 false, got ", searchFour)
	}
}
