package maxDepth

import (
	"algo/tree/tree"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	a := []int{3, 9, 20, 0, 0, 15, 7}
	root := tree.CreateTree(a)
	t.Log(MaxDepth(root))
}
