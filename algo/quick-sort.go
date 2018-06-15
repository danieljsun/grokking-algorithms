package algo

import (
	//"fmt"
	"math/rand"
)

func QuickSort(s []int) []int {
	/*
	if len <= 1
		reutrn s
	
	pivot = s[rand(len(s))]
	left = s where x < pivot
	right = s where x > pivot
	return left + x + right

	*/
	if len(s) <= 1 {
		return s
	}
	pivot := s[rand.Intn(len(s))]
	left := make([]int, 0, len(s))
	right := make([]int, 0, len(s))
	for i, _ := range s {
		if s[i] < pivot {
			left = append(left, s[i])
		}
		if s[i] > pivot {
			right = append(right, s[i])
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
