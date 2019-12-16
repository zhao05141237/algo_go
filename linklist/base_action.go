package main

import (
	"algo/linklist/singleLinkedList"
	"fmt"
)

func main() {
	//for i := 0; i <= 10; i++ {
	//	list.addNodeBefore(i)
	//}
	//node := list.addNodeAfter(0, list.head)
	//for i := 1; i <= 10; i++ {
	//	node = list.addNodeAfter(i, node)
	//}
	list := singleLinkedList.CreateListByLength(10)
	fmt.Printf("list length is:%d\n", list.GetLength())
	list.PrintList()
	//index := list.deleteNode(10)
	//index := list.FindNodeByData(10)
	//if index == list.GetLength() + 1 {
	//	fmt.Printf("not found you delelte node \n")
	//} else {
	//	fmt.Printf("found node and delete index :%d\n",index)
	//}
	//fmt.Printf("list length is:%d\n", list.GetLength())
	//list.PrintList()
}


