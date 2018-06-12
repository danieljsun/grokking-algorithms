package algo

func BinarySearch(s []int, n int) (found bool, index int) {
	low := 0
	high := len(s) - 1
	for !found && low <= high {
		mid := (low + high) / 2
		switch {
		case s[mid] == n:
			found = true
			index = mid
		case s[mid] < n:
			low = mid + 1
		case s[mid] > n:
			high = mid - 1
		}
	}
	return
}
