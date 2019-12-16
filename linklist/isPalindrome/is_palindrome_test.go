package isPalindrome

import (
	"algo/linklist/singleLinkedListNoHead"
	"fmt"
	"math"
	"testing"
)

/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
func isPalindrome(head *singleLinkedListNoHead.ListNode) bool {
	phead := &singleLinkedListNoHead.ListNode{
		Val:  0,
		Next: head,
	}
	slow := phead.Next

	//空
	if slow == nil {
		return true
	}

	//只有一个节点
	if slow.Next == nil {
		return true
	}

	fast := slow.Next.Next
	start := phead

	//只有两个结点
	if fast == nil {
		return slow.Val == slow.Next.Val
	}

	//三个结点
	if fast.Next == nil {
		return slow.Val == fast.Val
	}

	flagLeft := true
	middle := slow

	for ; slow != nil ; slow = slow.Next{
		if fast == nil || fast.Next == nil {
			if flagLeft && fast != nil {
				flagLeft = false
				continue
			}
			start = start.Next
			if slow.Val != start.Val {
				return false
			}
			flagLeft = false
		} else if slow == middle{
			start.Next = slow.Next
			slow.Next = slow.Next.Next
			start.Next.Next = slow
			fast = fast.Next.Next
		}else {
			middle.Next = middle.Next.Next
			slow.Next = start.Next
			start.Next = slow
			slow = middle
			fast = fast.Next.Next
		}

	}

	return true
}

/**
数学法 用时20ms 太慢 计算时间 都在power 涉及50000个节点的时候就会慢
 */
func isPalindromeMath(head *singleLinkedListNoHead.ListNode) bool {
	phead := &singleLinkedListNoHead.ListNode{
		Val:  0,
		Next: head,
	}
	slow := phead.Next

	//空
	if slow == nil {
		return true
	}

	//只有一个节点
	if slow.Next == nil {
		return true
	}

	fast := slow.Next.Next

	//只有两个结点
	if fast == nil {
		return slow.Val == slow.Next.Val
	}

	//三个结点
	if fast.Next == nil {
		return slow.Val == fast.Val
	}

	var sumLeft float64
	t := 0
	flagLeft := true

	//分左右两边之和
	for ; slow != nil ; slow = slow.Next{
		if fast == nil || fast.Next == nil {
			if flagLeft {
				sumLeft += math.Pow(10,float64(t)) * float64(slow.Val)
				if fast != nil {
					slow = slow.Next
				}
			}else {
				sumLeft -= math.Pow(10,float64(t)) * float64(slow.Val)
				t--
			}
			flagLeft = false
		} else {
			sumLeft += math.Pow(10,float64(t)) * float64(slow.Val)
			fast = fast.Next.Next
			t++
		}
	}

	return sumLeft == 0
}

func run(a []int) {
	list := &singleLinkedListNoHead.ListNode{}
	list = list.CreateList(a)
	list.PrintList(list)
	fmt.Println(isPalindrome(list))
}

func TestFirst(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	run(a)
}

func TestForth(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	run(a)
}

func TestSec(t *testing.T) {
	a := []int{1, 2, 3, 3, 2, 1}
	run(a)
}

func TestThird(t *testing.T) {
	a := []int{1, 1, 2, 1, 1}
	run(a)
}

func TestFifth(t *testing.T) {
	a := []int{1,2,2,2,1}
	run(a)
}

func TestSixth(t *testing.T)  {
	a := []int{1,3,2,4,3,2,1}
	run(a)
}

func TestSeventh(t *testing.T)  {
	a := []int{1,4,-1,-1,4,1}
	run(a)
}
