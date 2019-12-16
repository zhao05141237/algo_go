package removeElements

/**
删除链表中等于给定值 val 的所有节点。
示例:
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
 */
import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func removeElements(head *singleLinkedListNoHead.ListNode, val int) *singleLinkedListNoHead.ListNode {
	phead := &singleLinkedListNoHead.ListNode{
		Val:  0,
		Next: head,
	}

	if head == nil {
		return head
	}

	nNode := head.Next
	pNode := head

	if pNode.Val == val {
		phead.Next = nNode
		pNode = phead
	}

	for ;nNode != nil;nNode = nNode.Next{
		if nNode.Val == val {
			pNode.Next = nNode.Next
		} else {
			pNode = pNode.Next
		}

	}

	return phead.Next
}

func run(a []int, val int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	//list = middleNode(list)
	list = removeElements(list, val)
	list.PrintList(list)
}

func TestFirst(t *testing.T) {
	a := []int{1, 2, 6, 3, 4, 5, 6}
	val := 6
	run(a, val)
}

func TestSec(t *testing.T) {
	a := []int{1, 1, 2, 3, 3}
	val := 3
	run(a, val)
}

func TestThird(t *testing.T)  {
	a :=[]int{1}
	run(a,1)
}

func TestForth(t *testing.T)  {
	a :=[]int{1,1}
	run(a,1)
}
