package rotateRight

import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func rotateRight(head *singleLinkedListNoHead.ListNode, k int) *singleLinkedListNoHead.ListNode {

}

func run(a []int, k int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	list = rotateRight(list, k)
	list.PrintList(list)
}

func TestFirst(t *testing.T) {
	a := []int{1, 2, 3, 4}
	k := 4
	run(a, k)
}
