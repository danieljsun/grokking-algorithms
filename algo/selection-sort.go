package algo

func SelectionSort(s []int) []int {
	for i, _ := range s {
		for j := i + 1; j < len(s); j++ {
			if (s[i] > s[j]) {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
