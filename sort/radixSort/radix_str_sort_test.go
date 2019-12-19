package radixSort

import "testing"

func TestRadixStrSort(t *testing.T) {
	a := []string{"hke", "iba", "hzg", "ikf", "hac", "abcde"}
	RadixStrSort(a)
	t.Log(a)
}
