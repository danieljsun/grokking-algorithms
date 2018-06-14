package algo

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	test := func(input uint, expected uint) {
		fmt.Println("Factorial ", input)
		actual := Factorial(input)
		if actual != expected {
			t.Error("Expected ", expected, ", got ", actual)
		}
	}

	test(0, 1)
	test(1, 1)
	test(2, 2)
	test(3, 6)
}
