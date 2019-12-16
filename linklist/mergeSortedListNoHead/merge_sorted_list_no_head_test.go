package mergeSortedListNoHead

import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func mergeTwoListsSec(l1 *singleLinkedListNoHead.ListNode, l2 *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode   {
	//二者之一有空的 则返回另一个
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head1 := l1
	head2 := l2
	node1 := l1
	var nodeSec *singleLinkedListNoHead.ListNode

	for  node1 != nil && head2 != nil {
		if node1.Val >= head2.Val {
			//列表2 移除该节点
			nodeSec = head2
			head2 = nodeSec.Next
			//插入列表1node1节点前
			if head1 == node1 {
				nodeSec.Next = head1
				l1 = nodeSec
			} else {
				nodeSec.Next = node1
				head1.Next = nodeSec
			}
			head1 = nodeSec
		} else if node1.Val < head2.Val && node1.Next != nil {
			head1 = node1
			node1 = node1.Next
		} else if node1.Val < head2.Val && node1.Next == nil {
			node1.Next = head2
			head2 = head2.Next
			return l1
		}
	}

	return l1
}

func mergeTwoLists(l1 *singleLinkedListNoHead.ListNode, l2 *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
	//二者之一有空的 则返回另一个
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head1 := l1
	head2 := l2
	node1 := l1.Next
	node2 := l2.Next
	var nodeSec *singleLinkedListNoHead.ListNode

	if node1 == nil || node2 == nil {
		if node1 == nil && node2 == nil {
			if head1.Val > head2.Val {
				head2.Next = head1
				return l2
			} else {
				head1.Next = head2
				return l1
			}
		} else if node1 == nil && node2 != nil {
			if head1.Val <= head2.Val {
				head1.Next = head2
				l2 = head1
				return l2
			}
			for ; node2 != nil; node2 = node2.Next {
				if head1.Val <= node2.Val {
					head1.Next = node2
					head2.Next = head1
					return l2
				}
				head2 = node2
			}
			head2.Next = head1
			return l2
		} else if node2 == nil && node1 != nil {
			if head2.Val <= head1.Val {
				head2.Next = head1
				l1 = head2
				return l1
			}
			for ; node1 != nil; node1 = node1.Next {
				if head2.Val <= node1.Val {
					head2.Next = node1
					head1.Next = head2
					return l1
				}
				head1 = node1
			}
			head1.Next = head2
			return l1
		}

	}

	if head1.Val >= head2.Val {
		head2.Next = head1
		l1 = head2
		head2 = node2
		head1 = l1
		node1 = head1.Next
	}else {
		node1 = head1
	}

	for node1 != nil && head2 != nil {
		if node1.Val >= head2.Val {
			//列表2 移除该节点
			nodeSec = head2
			head2 = nodeSec.Next
			//插入列表1node1节点前
			nodeSec.Next = node1
			head1.Next = nodeSec
			head1 = nodeSec
		} else if node1.Val < head2.Val && node1.Next != nil {
			head1 = node1
			node1 = node1.Next
		} else if node1.Val < head2.Val && node1.Next == nil {
			node1.Next = head2
			head2 = head2.Next
			return l1
		}
	}

	if head2 != nil {
		head1.Next = head2
	}

	//for head2 != nil {
	//	nodeSec = head2
	//	head2 = nodeSec.Next
	//	//插入列表1后
	//	head1.Next = nodeSec
	//	head1 = head1.Next
	//}
	//l2.Next = nil
	return l1
}



func run(a1 []int,a2 []int) {
	l1 := &singleLinkedListNoHead.ListNode{}
	l2 := &singleLinkedListNoHead.ListNode{}

	l1 = l1.CreateList(a1)
	l2 = l2.CreateList(a2)

	l1.PrintList(l1)
	l2.PrintList(l2)

	//l3 := mergeTwoLists(l1, l2)
	l3 := mergeTwoListsSec(l1, l2)

	l3.PrintList(l3)

}

func TestFirst(t *testing.T)  {
	a1 := []int{-9,3}
	a2 := []int{5,7}
	run(a1,a2)
}

func TestSec(t *testing.T)  {
	a1 := []int{9,10}
	a2 := []int{1,100,200,1000}
	run(a1,a2)
}

func TestThird(t *testing.T)  {
	a1 := []int{1,2,4}
	a2 := []int{1,3,4}
	run(a1,a2)
}

func TestForth(t *testing.T)  {
	a1 := []int{-2,5}
	a2 := []int{-9,-6,-3,-1,1,6}
	run(a1,a2)
}

func TestFifth(t *testing.T)  {
	a1 := []int{1}
	a2 := []int{9}
	run(a1,a2)
}

func TestSixth(t *testing.T)  {
	a1 := []int{9}
	a2 := []int{1}
	run(a1,a2)
}

func TestSeventh(t *testing.T)  {
	a1 := []int{5}
	a2 := []int{1,2,4}
	run(a1,a2)
}
