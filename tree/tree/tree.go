package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//[3,9,20,null,null,15,7]
func CreateTree(a []int) *Node {
	if len(a) < 1 {
		return nil
	}
	root := &Node{
		Val:   a[0],
		Left:  createSubTree(a, 1),
		Right: createSubTree(a, 2),
	}

	return root
}

func createSubTree(a []int, index int) *Node {
	if index >= len(a) || a[index] <= 0 {
		return nil
	}
	return &Node{
		Val:   a[index],
		Left:  createSubTree(a, 2*index+1),
		Right: createSubTree(a, 2*index+2),
	}
}

func PreOrder(root *Node) {
	if root == nil || root.Val <= 0 {
		fmt.Println("tree is empty")
		return
	}
	fmt.Printf("%d", root.Val)
	preSubtreeOrder(root.Left)
	preSubtreeOrder(root.Right)
}

func preSubtreeOrder(root *Node) {
	if root != nil && root.Val > 0 {
		fmt.Printf("->%d", root.Val)
		preSubtreeOrder(root.Left)
		preSubtreeOrder(root.Right)
	} else {
		return
	}
}

func MiddleOrder(root *Node) {
	if root == nil || root.Val <= 0 {
		fmt.Println("tree is empty")
		return
	}
	middleSubtreeOrder(root)
}

func middleSubtreeOrder(root *Node) {
	if root != nil && root.Val > 0 {
		middleSubtreeOrder(root.Left)
		fmt.Printf("->%d", root.Val)
		middleSubtreeOrder(root.Right)
	} else {
		return
	}
}
