package algo

import "fmt"

func WhatSort(s []int) []int {
	for i, _ := range s {
		for j := i + 1; j < len(s); j++ {
			if (s[i] > s[j]) {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

// Returns the index of the smallest number in the array
func findSmallest(s []int) (smallestIndex int) {
	smallest := s[smallestIndex]
	for i := 1; i < len(s); i++ {
		if s[i] < smallest {
			smallestIndex = i
			smallest = s[i]
		}
	}
	fmt.Println("smallestIndex ", smallestIndex, ", smallest ", smallest)
	return
}

func SelectionSort(s []int) []int {
	for i, _ := range s {
		fmt.Println("Working on ", i);
		smallestIndex := findSmallest(s[i:])
		if smallestIndex > 0 {
			s[i], s[i + smallestIndex] = s[i + smallestIndex], s[i]
		}
		fmt.Println("After worked on ", i, " ", s);
	}
	return s
}
