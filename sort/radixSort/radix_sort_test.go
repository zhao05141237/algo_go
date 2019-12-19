package radixSort

import "testing"

func TestRadixSort(t *testing.T) {
	a := []int{430, 122, 332, 167, 899, 998, 455, 691, 571, 90, 9}
	t.Log(a)
	RadixSort(a)
	t.Log(a)
	t.Log(len("13305281715"))
}
