package mergeSortedListHead

import (
	"algo/linklist/singleLinkedList"
	"testing"
)

func mergeSortedList(list1 singleLinkedList.SingleLinkedList,list2 singleLinkedList.SingleLinkedList) singleLinkedList.SingleLinkedList{
	head1 := list1.GetHead()
	head2 := list2.GetHead()

	node1 := head1.GetNext()
	node2 := head2.GetNext()
	var nodeSec *singleLinkedList.Node

	//二者之一有空的 则返回另一个
	if node1 == nil {
		return list2
	}
	if node2 == nil {
		return list1
	}

	for node1 != nil && node2 != nil{
		if node1.GetData() >= node2.GetData() {
			//列表2 移除该节点
			nodeSec = node2
			node2 = node2.GetNext()
			head2.SetNext(node2)
			head2 = head2.GetNext()
			//插入列表1node1节点前
			nodeSec.SetNext(node1)
			head1.SetNext(nodeSec)
			head1 = nodeSec
		}else if node1.GetData() < node2.GetData() && node1.GetNext() != nil {
			head1 = node1
			node1 = node1.GetNext()
		}else if node1.GetData() < node2.GetData() && node1.GetNext() == nil {
			node1.SetNext(node2)
			return list1
		}

	}

	for node2 != nil {
		nodeSec = node2
		node2 = node2.GetNext()
		head2.SetNext(node2)
		head2 = head2.GetNext()
		//插入列表1后
		head1.SetNext(nodeSec)
		head1 = head1.GetNext()
	}
	list2.GetHead().SetNext(nil)
	return list1
}

func createList(list *singleLinkedList.SingleLinkedList,a []int) {
	var node *singleLinkedList.Node
	for index,val := range a {
		if index == 0 {
			node = list.AddNodeAfter(val,list.GetHead())
		} else {
			node = list.AddNodeAfter(val,node)
		}
	}
}

func run(a1 []int,a2 []int) {
	list1 := singleLinkedList.InitList()
	list2 := singleLinkedList.InitList()
	createList(&list1,a1)
	createList(&list2,a2)
	list1.PrintList()
	list2.PrintList()
	list3 := mergeSortedList(list1,list2)
	list3.PrintList()
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