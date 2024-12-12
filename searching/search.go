package searching

func LinearSearch(arr []int, val int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, val int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
