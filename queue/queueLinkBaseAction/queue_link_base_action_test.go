package queueLinkBaseAction

import (
	"algo/queue/queueLink"
	"testing"
)

func TestBaseAction(t *testing.T) {
	queueLinkList := queueLink.NewLinkedListQueue()
	for i := 1; i < 100; i++ {
		queueLinkList.EnQueue(i)
	}
GO1:
	for {
		v := queueLinkList.DeQueue()
		switch v.(type) {
		case bool:
			break GO1
		default:
			t.Log(v)
		}
	}

	t.Log(queueLinkList)
}
