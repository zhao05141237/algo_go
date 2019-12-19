package binarySearch

func BinarySearch(a []int, k int) int {
	if len(a) < 1 {
		return -1
	} else if len(a) == 1 {
		if a[0] == k {
			return 0
		} else {
			return -1
		}
	}

	low := 0
	high := len(a) - 1
	mid := low + (high-low)>>1

	for low <= high {
		if a[mid] < k {
			low = mid + 1
		} else if a[mid] > k {
			high = mid - 1
		} else {
			return mid + 1
		}
		mid = low + (high-low)>>1
	}

	return -1
}

func SearchRecursion(a []int, k int) int {
	if len(a) < 1 {
		return -1
	} else if len(a) == 1 {
		if a[0] == k {
			return 0
		} else {
			return -1
		}
	}
	return binarySearchRecursionC(a, 0, len(a)-1, (0+len(a)-1)/2, k)
}

func binarySearchRecursionC(a []int, low int, high int, mid int, k int) int {
	if low > high {
		return -1
	}

	if a[mid] == k {
		return mid + 1
	} else if a[mid] < k {
		low = mid + 1
	} else if a[mid] > k {
		high = mid - 1
	}
	mid = (low + high) / 2

	return binarySearchRecursionC(a, low, high, mid, k)
}
