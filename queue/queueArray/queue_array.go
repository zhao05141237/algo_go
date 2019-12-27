package queueArray

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (a *ArrayQueue) EnQueue(v interface{}) bool {
	if a.tail == a.capacity {
		//queue is full
		if a.head == 0 {
			return false
		}
		//队列满了 再搬移出队数据 不用每次出队都搬移数据
		for i := a.head; i < a.capacity; i++ {
			a.q[i-a.head] = a.q[i]
		}
		a.tail -= a.head
		a.head = 0
	}
	a.q[a.tail] = v
	a.tail++
	return true
}

func (a *ArrayQueue) DeQueue() interface{} {
	if a.head == a.tail {
		//队列为空
		return false
	}
	v := a.q[a.head]
	a.head++
	return v
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.head == a.tail
}

func (a *ArrayQueue) String() string {
	if a.head == a.tail {
		return "empty queue"
	}
	result := "head"
	for i := a.head; i < a.tail; i++ {
		result += fmt.Sprintf("<-%+v", a.q[i])
	}
	result += "<-tail"
	return result
}
