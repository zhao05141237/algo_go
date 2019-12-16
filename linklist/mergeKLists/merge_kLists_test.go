package mergeKLists

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
import (
	"algo/linklist/singleLinkedListNoHead"
	"testing"
)

func mergeKLists(lists []*singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
	if len(lists) <= 1 {
		return lists[0]
	}
	listsTmp := mergeTwoListsSec(lists[0], lists[1])

	for i := 2; i < len(lists); i++ {
		listsTmp = mergeTwoListsSec(listsTmp, lists[i])
	}

	return listsTmp
}

func mergeTwoListsSec(l1 *singleLinkedListNoHead.ListNode, l2 *singleLinkedListNoHead.ListNode) *singleLinkedListNoHead.ListNode {
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

	for node1 != nil && head2 != nil {
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

func run(a [][]int) {
	lists := make([]*singleLinkedListNoHead.ListNode, len(a))
	for i := 0; i < len(a); i++ {
		lists[i] = (&singleLinkedListNoHead.ListNode{}).CreateList(a[i])
	}
	for i := 0; i < len(lists); i++ {
		lists[i].PrintList(lists[i])
	}
	result := mergeKLists(lists)
	result.PrintList(result)
}

func TestFirst(t *testing.T) {
	a := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	run(a)
}
