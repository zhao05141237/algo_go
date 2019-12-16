package main

import (
	"algo/linklist/singleLinkedList"
	"errors"
	"flag"
	"fmt"
)

func inverseList(list *singleLinkedList.SingleLinkedList) {
	if list.IsEmpty() {
		return
	}
	tail := list.GetHead()
	var p *singleLinkedList.Node
	//tail指向最后一个节点
	for ; tail.GetNext() != nil; tail = tail.GetNext() {
	}
	//遍历list 直到header.next指向tail为止
	for list.GetHead().GetNext() != tail {
		//将header->next节点 移除
		p = list.GetHead().GetNext()
		list.GetHead().SetNext(p.GetNext())
		if tail.GetNext() == nil {
			//tail是最后一个节点  插入并设置为最后一个节点
			p.SetNext(nil)
			tail.SetNext(p)
		} else {
			//tail和尾结点中间插入
			p.SetNext(tail.GetNext())
			tail.SetNext(p)
		}
	}
}

var length int

func init() {
	flag.IntVar(&length, "length", 10, "list length")
}

func main() {
	flag.Parse()
	if length <= 0 {
		panic(errors.New("list length gt zero!"))
	}
	list := singleLinkedList.CreateListByLength(length)
	list.PrintList()
	fmt.Println("after inverse")
	inverseList(&list)
	list.PrintList()
}
