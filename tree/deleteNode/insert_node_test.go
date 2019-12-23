package deleteNode

import (
	"algo/tree/tree"
	"testing"
)

func TestInsertNode(t *testing.T) {
	//a := []int{5, 3, 6, 2, 4, 0, 7}
	//a := []int{3,1,4,0,2}
	//a := []int{3,2,4,1}
	a := []int{15, 3, 16, 1, 4, 0, 0, 0, 2}
	root := tree.CreateTree(a)
	_ = InsertNode(root, 5)
}
