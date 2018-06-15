package algo

func binarySearchRecursive(s []int, n int) (found bool, index int) {
	/*
	if len is 0
		return false, 0
	mid = start + end)/ 2
	if s[mid] equals n
		return true, mid
	if n < s[mid]
		return rec(s[:mid, n)
	if n > s[mid]
		return mid + rec(s[mid+1:], n)
	return false
	else
	*/
	low := 0
	high := len(s) - 1
	if len(s) > 0 {
		mid := (low + high) / 2
		if s[mid] == n {
			return true, mid
		}

		if n < s[mid] {
			return binarySearchRecursive(s[:mid], n)
		}

		if s[mid] < n {
			found, index = binarySearchRecursive(s[mid+1:], n)
			if found {
				return found, mid + index
			}
			return found, index
		}
	}

	return
}

func binarySearchNonRecursive(s []int, n int) (found bool, index int) {
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

func BinarySearch(s []int, n int) (found bool, index int) {
	return binarySearchRecursive(s, n)
}
