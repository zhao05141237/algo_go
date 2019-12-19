package queueLink

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type QueueLink struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedListQueue() *QueueLink {
	return &QueueLink{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *QueueLink) EnQueue(v interface{}) {
	node := &Node{
		data: v,
		next: nil,
	}
	if l.head == nil {
		//空队列
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

func (l *QueueLink) DeQueue() interface{} {
	if l.head == nil {
		//队列空
		return false
	}
	v := l.head.data
	l.head = l.head.next
	l.length--
	return v
}

func (l *QueueLink) String() string {
	if l.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := l.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.data)
	}
	result += "<-tail"
	return result
}
