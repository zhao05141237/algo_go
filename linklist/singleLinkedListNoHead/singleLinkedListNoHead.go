package singleLinkedListNoHead

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) CreateNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (node *ListNode) InsertIntoList(a int) *ListNode {
	tmpNode := node.CreateNode(a)
	if node == nil {
		return tmpNode
	}
	p := node
	for p.Next != nil {
		p = p.Next
	}
	p.Next = tmpNode
	return node
}

func (node *ListNode) FindValInList(a int) bool {
	if node == nil {
		return false
	}
	p := node
	for p != nil {
		if p.Val == a {
			return true
		}
		p = p.Next
	}
	return false
}

func (node *ListNode) CreateList(a []int) *ListNode {
	var l *ListNode
	var p *ListNode

	for index, val := range a {
		tmpNode := node.CreateNode(val)
		if index == 0 {
			l = tmpNode
			p = l
		} else {
			p.Next = tmpNode
			p = p.Next
		}
	}
	return l
}

func (node *ListNode) PrintList(l *ListNode) {
	p := l
	for ; p != nil; p = p.Next {
		if p.Next != nil {
			fmt.Printf("%d->", p.Val)
		} else {
			fmt.Printf("%d\n", p.Val)
		}
	}
	return
}
