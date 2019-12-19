package queueCircleAction

import (
	"algo/queue/queueCircle"
	"testing"
)

func TestAction(t *testing.T) {
	queue := queueCircle.NewCircleQueue(10)

	for i := 1; i < 100; i++ {
		if !queue.EnQueue(i) {
			break
		}
	}
	t.Log(queue)

GO1:
	for i := 1; i < 100; i++ {
		v := queue.DeQueue()
		switch v.(type) {
		case bool:
			break GO1
		default:
			t.Log(v)
		}
	}
	t.Log(queue)

	for i := 1; i < 100; i++ {
		if !queue.EnQueue(i) {
			break
		}
	}

	t.Log(queue)

}
