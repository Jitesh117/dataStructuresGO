package sorting

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pIndex := partition(arr, low, high)

		quickSortHelper(arr, low, pIndex-1)
		quickSortHelper(arr, pIndex+1, high)
	}
}

func QuickSort(arr []int) []int {
	quickSortHelper(arr, 0, len(arr)-1)
	return arr
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
