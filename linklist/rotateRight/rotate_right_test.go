package rotateRight

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/
import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func rotateRight(head *singleLinkedListNoHead.ListNode, k int) *singleLinkedListNoHead.ListNode {
	phead := &singleLinkedListNoHead.ListNode{
		Val:  0,
		Next: head,
	}

	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	length := 0
	node := phead.Next
	for ; node != nil; node = node.Next {
		length++
	}

	if k > length {
		k = k % length
	}

	i := 1
	for i <= k {
		start := phead
		node := phead.Next
		preNode := node
		prePreNode := node

		//移动node到末尾
		for ; node != nil; node = node.Next {
			prePreNode = preNode
			preNode = node
		}

		prePreNode.Next = preNode.Next
		preNode.Next = start.Next
		start.Next = preNode

		i++
	}

	return phead.Next
}

func run(a []int, k int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	list = rotateRight(list, k)
	list.PrintList(list)
}

func TestFirst(t *testing.T) {
	a := []int{0, 1, 2}
	k := 4
	run(a, k)
}
