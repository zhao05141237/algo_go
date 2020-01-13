package matrix

import "testing"

func TestMultiplication(t *testing.T) {
	a := [][]int{{1, 2, 3}, {3, 4, 5}, {4, 5, 6}}
	b := [][]int{{1, 2, 1}, {1, 3, 1}, {1, 5, 1}}

	t.Log(Multiplication(a, b))
}
