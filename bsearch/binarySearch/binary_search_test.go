package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	//a := []int{1,3,5,7,9,11,13}
	//t.Log(BinarySearch(a,5))
	//t.Log(BinarySearch(a,1))
	//t.Log(BinarySearch(a,13))
	//t.Log(BinarySearch(a,9))
	//t.Log(BinarySearch(a,10))
	//a = []int{}
	//t.Log(BinarySearch(a,0))
	//a = []int{1}
	//t.Log(BinarySearch(a,1))
	//t.Log(BinarySearch(a,2))

	a := []int{-1, 0, 3, 5, 9, 12}
	t.Log(BinarySearch(a, 9))
}

func TestBinarySearchRecursion(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 11, 13}
	t.Log(SearchRecursion(a, 5))
	t.Log(SearchRecursion(a, 1))
	t.Log(SearchRecursion(a, 13))
	t.Log(SearchRecursion(a, 9))
	t.Log(SearchRecursion(a, 10))
	a = []int{}
	t.Log(SearchRecursion(a, 0))
	a = []int{1}
	t.Log(SearchRecursion(a, 1))
	t.Log(SearchRecursion(a, 2))
}
