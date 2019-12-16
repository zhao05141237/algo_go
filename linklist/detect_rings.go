package main

import (
	"algo/linklist/singleLinkedList"
	"errors"
	"flag"
	"fmt"
)

func detectRings2(list *singleLinkedList.SingleLinkedList) bool{
	p := list.GetHead().GetNext()
	q := list.GetHead()

	for p != q {
		if p == nil || p.GetNext() == nil {
			return false
		}
		p = p.GetNext()
		p = p.GetNext()
		q = q.GetNext()
	}
	return true
}

func detectRings(list *singleLinkedList.SingleLinkedList) bool{
	tail := list.GetHead().GetNext()
    for ;tail != nil;tail = tail.GetNext() {
		if tail.GetNext() == tail {
			//自身结点循环
			return true
		}
		p := list.GetHead()
		for ; p != tail;p = p.GetNext() {
			if tail.GetNext() == p {
				//有环
				return true
			}
		}
	}
	return false
}

func creatRingsByPos(list *singleLinkedList.SingleLinkedList,pos int) {
	if pos < 0 || pos >= list.GetLength() {
		return
	}
	index :=0
	tail := list.GetHead().GetNext()
	q,p := list.GetHead(),list.GetHead()

	for ;tail!= nil;tail = tail.GetNext() {
		if index == pos {
			p = tail
		}
		q = tail
		index++
	}

	//构造一个环
	if  p != list.GetHead() {
		q.SetNext(p)
	}
}

var detectListLength int
var detectListPos int

func init() {
	flag.IntVar(&detectListLength, "length", 10, "list length")
	flag.IntVar(&detectListPos, "pos", -1, "tail node next position")
}

func main() {
	flag.Parse()
	if detectListLength <= 0 {
		panic(errors.New("list length gt zero!"))
	}
	list := singleLinkedList.CreateListByLength(detectListLength)
	creatRingsByPos(&list,detectListPos)
	//ringFlag := detectRings(&list)
	ringFlag := detectRings2(&list)
	fmt.Printf("detect rings resulst is :%v\n", ringFlag)
	if !ringFlag {
		list.PrintList()
	}
}
