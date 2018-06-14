package algo

func Factorial(n uint) uint {
	if n <= 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
