package string

import "testing"

func TestBfSearch(t *testing.T) {
	t.Log(bfSearch("abcd227fac", "abcd"))
}

func TestRkHashValue(t *testing.T) {
	t.Log(hashValue("cba"))
}

func TestRkSearch(t *testing.T) {
	t.Log(rkSearch("abcdfac", "c"))
}
