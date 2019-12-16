package main

import (
	"fmt"
)

func insertSort(a []int)  {
	aLen := len(a)
	for i:=1; i < aLen; i++ {
		k := a[i]
		t := i
		for j:=i-1; j >= 0;j-- {
			if a[j] > k {
				a[j+1] = a[j]
				t = j
			}
		}
		if t != i {
			a[t] = k
		}
	}
}

func insertSort2(a []int)  {
	aLen := len(a)
	for i:=1; i < aLen; i++ {
		k := a[i]
		j:=i-1
		for ; j >= 0;j-- {
			if a[j] > k {
				a[j+1] = a[j]
			}else {
				break
			}
		}
		a[j+1] = k
	}
}

func bubbleSort(a []int)  {
	aLen := len(a)
	for i:=0;i<aLen;i++ {
		flag := true
		for j:=0;j<aLen-i-1;j++ {
			if a[j] > a[j+1]  {
				a[j],a[j+1] = a[j+1],a[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func selectSort(a []int)  {
	aLen := len(a)
	for i:= 0; i < aLen; i++ {
		k,min := i,a[i]
		for j:=i+1;j<aLen;j++ {
			if a[j] < min {
				min = a[j]
				k = j
			}
		}
		if k != i {
			a[i],a[k] = a[k],a[i]
		}
	}
}

func main() {
	a := []int{3,5,4,1,2,6}
	fmt.Println(a)
	//bubbleSort(a)
	//insertSort(a)
	//insertSort2(a)
	selectSort(a)
	fmt.Println(a)
}
