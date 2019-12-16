package removeNthFromEnd

import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？
 */

func removeNthFromEnd(head *singleLinkedListNoHead.ListNode, n int) *singleLinkedListNoHead.ListNode {
	if head == nil || n < 1 {
		return head
	}
	//只有一个节点情况
	if head.Next == nil {
		if n == 1 {
			//是1 则删除当前节点
			head = nil
			return head
		} else {
			return head
		}
	}
	length := 0
	node := head
	for ;node != nil;node = node.Next {
		length++
	}
	index := length - n + 1
	//n 大于链表长度无效
	if index <= 0 {
		return head
	}else if index == 1 {
		//n正好等于链表长度
		//删除第一个节点
		head = head.Next
		return head
	}

	p := head
	deleteIndex := 1
	node = head //p node 都指向head
	for ;node != nil;node = node.Next {
		if deleteIndex == index {
			//删除
			p.Next = node.Next
			return head
		}else{
			p = node
		}
		deleteIndex++
	}
	return head
}

func run(a []int,n int){
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	list = removeNthFromEnd(list,n)
	list.PrintList(list)
}

func TestFirst(t *testing.T)  {
	a := []int{1,2,3,4,5}
	n := 2
	run(a,n)
}

func TestSec(t *testing.T)  {
	a := []int{1,2,3,4,5}
	n := 9
	run(a,n)
}

func TestThird(t *testing.T)  {
	a := []int{1,2,3,4,5}
	n := -1
	run(a,n)
}

func TestForth(t *testing.T)  {
	a := []int{1,2,3,4,5}
	n := 5
	run(a,n)
}
