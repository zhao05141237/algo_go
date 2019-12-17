package stackArithmetic

import (
	"algo/stack/stackList"
	"strconv"
	"testing"
)

func createStack(n int) *stackList.StackList {
	stack := &stackList.StackList{}
	stack.CreateStack(n)
	return stack
}

func pushStack(list *stackList.StackList, v interface{}) {
	switch v.(type) {
	case int:
		_, _ = list.Push(strconv.Itoa(v.(int)))
	default:
		_, _ = list.Push(v.(string))
	}
}

func popStack(list *stackList.StackList) string {
	elem, _ := list.Pop()
	return elem
}

func topStack(list *stackList.StackList) string {
	elem, _ := list.Top()
	return elem
}

func compareTopAndOpLevel(list *stackList.StackList, op string) bool {
	topLevel := operationLevel(topStack(list))
	opLevel := operationLevel(op)
	return topLevel < opLevel
}

func computeResultAndPushNumberStack(numberStack *stackList.StackList, operationStack *stackList.StackList) {
	var opVal string
	var numberValLeft, numberValRight, result int
	opVal = popStack(operationStack)
	numberValRight, _ = strconv.Atoi(popStack(numberStack))
	numberValLeft, _ = strconv.Atoi(popStack(numberStack))
	result = operationCompute(numberValLeft, numberValRight, opVal)
	pushStack(numberStack, result)
}

func TestSec(t *testing.T) {
	express := "3+5*8-6"
	t.Log(len(express))
	for index := range express {
		t.Log(index)
	}
}

func TestArithmetic(t *testing.T) {
	express := "1*2*3"
	result := 0
	number := 0
	operationStack := createStack(10)
	numberStack := createStack(10)
	expressLength := len(express)

	for index, val := range express {
		op := string(val)
		if isOperation(op) {
			pushStack(numberStack, number)
			if operationStack.IsEmpty() {
				pushStack(operationStack, op)
			} else {
				if compareTopAndOpLevel(operationStack, op) {
					pushStack(operationStack, op)
				} else {
					for {
						if compareTopAndOpLevel(operationStack, op) {
							break
						}
						computeResultAndPushNumberStack(numberStack, operationStack)
						if operationStack.IsEmpty() {
							pushStack(operationStack, op)
							break
						}
					}
				}
			}
			number = 0
		} else if isDigit(op) {
			value, _ := strconv.Atoi(op)
			number = number*10 + value
		}

		if index == expressLength-1 {
			pushStack(numberStack, number)
			for {
				computeResultAndPushNumberStack(numberStack, operationStack)
				if operationStack.IsEmpty() {
					break
				}
			}
		}
	}
	if !numberStack.IsEmpty() {
		result, _ = strconv.Atoi(popStack(numberStack))
	}
	t.Logf("%s=%d", express, result)
}

func isDigit(str string) bool {
	return str >= "0" && str <= "9"
}

func isOperation(str string) bool {
	return str == "+" || str == "-" || str == "*" || str == "/"
}

func operationLevel(str string) int {
	if str == "+" || str == "-" {
		return 1
	} else if str == "*" || str == "/" {
		return 2
	} else {
		return 0
	}
}

func operationCompute(left int, right int, op string) int {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		return 0
	}
}
