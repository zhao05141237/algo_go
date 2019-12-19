package quickSort

func QuickSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}
	quickSortC(a, 0, arrLen-1)
}

func quickSortC(a []int, start int, end int) {
	if start >= end {
		return
	}
	//找一个中间点 middle
	p := partition(a, start, end)
	quickSortC(a, start, p-1)
	quickSortC(a, p+1, end)
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			//swap arr[i] arr[j]
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	//swap arr[i] arr[end] i is partition
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
