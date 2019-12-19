package queueCircle

import "fmt"

type CircleQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircleQueue(n int) *CircleQueue {
	return &CircleQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (a *CircleQueue) EnQueue(v interface{}) bool {
	if (a.tail+1)%a.capacity == a.head {
		//queue is full
		return false
	}
	a.tail = (a.tail + 1) % a.capacity
	a.q[a.tail] = v
	return true
}

func (a *CircleQueue) DeQueue() interface{} {
	if a.head == a.tail {
		//队列为空
		return false
	}
	a.head = (a.head + 1) % a.capacity
	v := a.q[a.head]
	return v
}

func (a *CircleQueue) String() string {
	if a.head == a.tail {
		return "empty queue"
	}
	result := "head"
	for i := a.head + 1; i%a.capacity != a.tail; i++ {
		result += fmt.Sprintf("<-%+v", a.q[i%a.capacity])
	}
	result += "<-tail"
	return result
}
