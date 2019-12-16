package reverseBetween

import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

/**
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。
示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
 */

/**
遍历且计node节点数 [m,n]区间节点一一反转
 */
func reverseBetween(head *singleLinkedListNoHead.ListNode, m int, n int) *singleLinkedListNoHead.ListNode {
	plead := head.CreateNode(0)
	plead.Next = head
	start,node,middle := plead,head,head
	mNode,nNode := head,head
	indexNode := 0

	//m>=n 则不反转 返回原链表
	if m < n {
		for ;node != nil; node = node.Next {
			indexNode++
			if indexNode + 1 == m {
				start = node
				middle = node.Next
			}else if indexNode == m {
				mNode = node
				nNode = node.Next
				node.Next = nNode.Next
				nNode.Next = node
				start.Next = nNode
				indexNode++
			} else if indexNode > m  && indexNode <= n {
				mNode = node
				nNode = mNode.Next
				middle.Next = nNode
				mNode.Next = start.Next
				start.Next = mNode
				node = middle
			}else if indexNode > n {
				break
			}
		}
	}
	return plead.Next
}

func run(a []int,m int,n int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	list = reverseBetween(list,m,n)
	list.PrintList(list)
}

func TestFirst(t *testing.T){
	a:=[]int{1,2,3,4}
	m,n:=1,4
	//m,n := 4,5
	run(a,m,n)
}
