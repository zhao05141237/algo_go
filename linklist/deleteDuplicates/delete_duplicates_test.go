package deleteDuplicates

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */
import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func deleteDuplicates(head *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
	phead := &singleLinkedListNoHead.ListNode{
		Val:  0,
		Next: head,
	}

	if head == nil {
		return head
	}

	nNode := head.Next
	pNode := head

	for ; nNode != nil; nNode = nNode.Next {
		if nNode.Val == pNode.Val {
			pNode.Next = nNode.Next
		}else{
			pNode = pNode.Next
		}
	}

	return phead.Next
}

func run(a []int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	//list = middleNode(list)
	list = deleteDuplicates(list)
	list.PrintList(list)
}

func TestFirst(t *testing.T) {
	a := []int{1, 1, 2}
	run(a)
}

func TestSec(t *testing.T)  {
	a := []int{1,1,2,3,3}
	run(a)
}