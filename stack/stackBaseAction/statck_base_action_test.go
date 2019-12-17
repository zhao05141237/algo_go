package stackBaseAction

import (
	"algo/stack/stackList"
	"testing"
)

func createStack(n int) *stackList.StackList {
	stack := &stackList.StackList{}
	stack.CreateStack(n)
	return stack
}

func TestFirst(t *testing.T) {
	stack := createStack(7)
	elements := "hello world!"

	for _, val := range elements {
		if _, err := stack.Push(string(val)); err != nil {
			break
		}
	}

	//if elem,err := stack.Top();err != nil {
	//	t.Log(err)
	//}else{
	//	t.Log(elem)
	//}

	for {
		if elem, err := stack.Pop(); err != nil {
			t.Log(err)
			break
		} else {
			t.Log(elem)
		}

	}
}
