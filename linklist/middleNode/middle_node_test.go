package middleNode

import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func middleNode(head *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
	plead := head.CreateNode(0)
	plead.Next = head
	middle, length := 0, 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	if length <= 1 {
		//长度为1的链表没有中间结点
		return head
	} else {
		middle = length/2 + 1
	}
	length = 0
	for node := head; node != nil; node = node.Next {
		length++
		if length == middle {
			plead.Next = node
			break
		}
	}
	return plead.Next
}

//利用快慢指针，快指针每次走两步，慢指针每次走一步，所以快指针走的距离为慢指针的两倍，故当快指针遍历到链表末尾时，慢指针指向记为中间节点
func middleNodeWidthSlowFastPoint(head *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
	plead := &singleLinkedListNoHead.ListNode{
		Val:  0,
		Next: head,
	}
	slow,fast:= head,head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	plead.Next = slow
	return plead.Next
}

func run(a []int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	//list = middleNode(list)
	list = middleNodeWidthSlowFastPoint(list)
	list.PrintList(list)
}

func TestFirst(t *testing.T){
	a:=[]int{1,2,3,4,5}
	run(a)
}

func TestSec(t *testing.T){
	a:=[]int{1,2,3,4,5,6}
	run(a)
}

func TestThird(t *testing.T)  {
	a:=make([]int,0)
	run(a)
}

func TestForth(t *testing.T)  {
	a:=[]int{1}
	run(a)
}