package coutingSort

import "testing"

func TestCoutingSort(t *testing.T) {
	a := []int{2, 5, 3, 0, 2, 3, 0, 3}
	t.Log(a)
	CoutingSort(a)
	t.Log(a)
}
