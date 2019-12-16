package singleLinkedList

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data interface{}
	next *Node
}

func (node *Node) Node(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func (node *Node) GetNext() *Node {
	return node.next
}

func (node *Node) SetNext(next *Node)  {
	node.next = next
}

func (node *Node) GetData() int64 {
	if val,ok := (node.data).(int64);ok{
		return val
	} else {
		return 0
	}
}

type SingleLinkedList struct {
	head   *Node
	length int
}

func (list *SingleLinkedList) CreateList() {
	head := creatNode(nil)
	list.head = head
	list.length = 0
}

func (list *SingleLinkedList) AddNodeBefore(data interface{}) {
	node := creatNode(data)
	if list.length == 0 {
		list.head.next = node
	} else {
		node.next = list.head.next
		list.head.next = node
	}
	list.length++
}

func (list *SingleLinkedList) AddNodeAfter(data interface{}, before *Node) *Node {
	node := creatNode(data)
	before.next = node
	list.length++
	return node
}

func (list *SingleLinkedList) DeleteNode(data interface{}) int {
	q,p := list.head,list.head
	index := 0
	for ; p != nil; p = p.next {
		if p.data == data {
			node := p
			q.next = p.next
			node.next = nil
			list.length--
			break
		}
		q = p
		index++
	}
	return index
}

func (list *SingleLinkedList) FindNodeByData(data interface{}) int {
	p := list.head
	index := 0
	for ; p != nil; p = p.next {
		if p.data == data {
			break
		}
		index++
	}
	return index
}

func (list *SingleLinkedList) IsIntNodeDataType() bool {
	p := list.head
	p = p.next
	for ; p != nil; p = p.next {
		if _,ok := (p.data).(int64);!ok {
			return false
		}
	}
	return true
}

func (list *SingleLinkedList) GetLength() int {
	return list.length
}

func (list *SingleLinkedList) SetLength(length int) {
	list.length = length
}

func (list *SingleLinkedList) IsEmpty() bool {
	return list.length <= 0
}

func (list *SingleLinkedList) GetHead() * Node {
	return list.head
}

func (list *SingleLinkedList) PrintList() {
	p := list.head
	if p.next == nil {
		fmt.Println("empty list!")
		return
	}
	p = p.next
	for ; p != nil; p = p.next {
		fmt.Printf("%d", p.data)
		if p.next != nil {
			fmt.Printf("->")
		}
	}
	fmt.Printf("\n")
}

func creatNode(data interface{}) *Node {
	node := Node{}
	return node.Node(data)
}

func InitList() SingleLinkedList {
	list := SingleLinkedList{}
	list.CreateList()
	return list
}


func CreateListByLength(length int) SingleLinkedList {
	rand.Seed(time.Now().UnixNano())
	list := InitList()
	for i:=0; i<length; i++{
		list.AddNodeBefore(rand.Intn(100))
	}
	return list
}


