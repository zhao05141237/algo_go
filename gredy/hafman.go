package gredy

import (
	"algo/sort/quickSort"
	"algo/tree/tree"
	"fmt"
)

//a [2 3 4 4 5 7]
func BuildHafManTree(a []int) {
	mapTree := map[int]*tree.Node{}
	length := len(a)
	var t *tree.Node
	for length > 1 {
		fmt.Println(a)
		quickSort.QuickSort(a)
		first := a[0]
		second := a[1]
		sum := first + second
		t, mapTree = huffmanBuildTree(first, second, mapTree)
		a = a[2:]
		a = append(a, sum)
		length = length - 1
	}
	huffmanTreeOrder(t)
}

func huffmanBuildTree(first int, second int, mapTree map[int]*tree.Node) (*tree.Node, map[int]*tree.Node) {
	var left, right *tree.Node
	sum := first + second
	t := &tree.Node{
		Val:   sum,
		Left:  nil,
		Right: nil,
	}
	if val, ok := mapTree[first]; ok && val != nil {
		left = val
		mapTree[first] = nil
	} else {
		left = &tree.Node{
			Val:   first,
			Left:  nil,
			Right: nil,
		}
	}

	if val, ok := mapTree[second]; ok && val != nil {
		right = val
		mapTree[second] = nil
	} else {
		right = &tree.Node{
			Val:   second,
			Left:  nil,
			Right: nil,
		}
	}

	t.Left = left
	t.Right = right
	mapTree[sum] = t
	return t, mapTree
}

func huffmanTreeOrder(root *tree.Node) {
	if root == nil {
		return
	}
	huffmanTreeOrderC(root.Left, "0")
	huffmanTreeOrderC(root.Right, "1")
}

func huffmanTreeOrderC(root *tree.Node, code string) {
	if root.Left == nil && root.Right == nil {
		fmt.Printf("tree node %d code is :%s\n", root.Val, code)
		return
	}
	huffmanTreeOrderC(root.Left, code+"0")
	huffmanTreeOrderC(root.Right, code+"1")
}
