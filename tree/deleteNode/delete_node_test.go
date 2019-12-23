package deleteNode

import (
	"algo/tree/tree"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	a := []int{5, 3, 6, 2, 4, 0, 7}
	//a := []int{3,1,4,0,2}
	//a := []int{3,2,4,1}
	//a := []int{5,3,6,1,4,0,0,0,2}
	root := tree.CreateTree(a)
	_ = DeleteNode(root, 3)
}
