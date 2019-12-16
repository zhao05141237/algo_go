package main

import "fmt"

func quickSort(a []int) {
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

func partition(arr []int,start int,end int) int{
	pivot := arr[end]
	i := start
	for j:=start;j < end;j++ {
		if arr[j] < pivot {
			//swap arr[i] arr[j]
			arr[i],arr[j] = arr[j],arr[i]
			i++
		}
	}
	//swap arr[i] arr[end] i is partition
	arr[i],arr[end] = arr[end],arr[i]
	return i
}

//func partition(arr []int, start int, end int) int {
//	pivot := arr[end]
//	tmpLeft := make([]int,end-start+1)
//	tmpRight := make([]int,end-start+1)
//	i,j:=0,0
//	for k:=start; k <= end; k++ {
//		if arr[k] < pivot {
//			tmpLeft[i] = arr[k]
//			i++
//		} else if arr[k] > pivot{
//			tmpRight[j] = arr[k]
//			j++
//		}
//	}
//
//	for index:=0;index<i;index++{
//		arr[index+start] = tmpLeft[index]
//	}
//	pivotPosition := i+start
//	arr[pivotPosition] = pivot
//	for index:=0;index<j;index++{
//		arr[index+pivotPosition+1] = tmpRight[index]
//	}
//	return pivotPosition
//}

func main() {
	a := []int{11, 8, 3, 9, 7, 1, 2, 5, 13}
	fmt.Println(a)
	quickSort(a)
	fmt.Println(a)
}
