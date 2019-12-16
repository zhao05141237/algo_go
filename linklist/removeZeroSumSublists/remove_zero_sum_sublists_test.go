package removeZeroSumSublists

import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func removeZeroSumSublists(head *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
	plead := head.CreateNode(0)
	plead.Next = head

	q := plead
GO1:	for node := head; node != nil; node = node.Next {
		val := 0
		p := q
		for mnode := node; mnode != nil; mnode = mnode.Next {
			val += mnode.Val
			if val == 0 {
				p.Next = mnode.Next
				p = p.Next
				node.Next = p
				continue GO1
			}
		}
		q = q.Next
	}
	return plead.Next
}

func run(a []int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	list = removeZeroSumSublists(list)
	list.PrintList(list)
}

func TestFirst(t *testing.T) {
	a := []int{1,2,3,-3,-2}
	run(a)
}

func TestSec(t *testing.T) {
	a := []int{1,3,2,-3,-2,5,5,-5,1}
	run(a)
}
