package main

import (
	"algo/linklist/singleLinkedList"
	"flag"
	"strconv"
	"strings"
)

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */

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

var strList1 string
var strList2 string

func init() {
	flag.StringVar(&strList1,"list1","1,2,4","sorted list1 implode with ,")
	flag.StringVar(&strList2,"list2","1,3,4","sorted list2 implode with ,")
}

func createListByString(list *singleLinkedList.SingleLinkedList,strList string) {
	listArr := strings.Split(strList,",")
	var node *singleLinkedList.Node
	for index,val := range listArr {
		interval,_ := strconv.ParseInt(val,10,8)
		if index == 0 {
			node = list.AddNodeAfter(interval,list.GetHead())
		} else {
			node = list.AddNodeAfter(interval,node)
		}
	}
}

func main() {
	flag.Parse()
	list1 := singleLinkedList.InitList()
	list2 := singleLinkedList.InitList()
	createListByString(&list1,strList1)
	createListByString(&list2,strList2)
	list1.PrintList()
	list2.PrintList()
	list3 := mergeSortedList(list1,list2)
	list3.PrintList()
}
