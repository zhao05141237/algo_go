package main

import "fmt"

func mergeSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}
	mergesortC(a, 0, arrLen-1)
}

func mergesortC(a []int, start int, end int) {
	if start >= end {
		return
	}
	//找一个中间点 middle
	middle := (start + end) / 2
	mergesortC(a, start, middle)
	mergesortC(a, middle+1, end)
	merge(a, start, middle, end)
}

func merge(arr []int, start int, mid int, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}

func main() {
	a := []int{11, 8, 3, 9, 7, 1, 2, 5,13}
	fmt.Println(a)
	mergeSort(a)
	fmt.Println(a)
}
