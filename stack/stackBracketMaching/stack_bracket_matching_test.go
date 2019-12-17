package stackBracketMaching

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

func TestBracketMatching(t *testing.T) {
	express := "((1*2*3{["
	expressLength := len(express)
	var result bool
	stack := createStack(expressLength)

	for _, val := range express {
		backedVal := string(val)
		if isLeftBracket(backedVal) {
			pushStack(stack, backedVal)
		} else if isRightBracket(backedVal) {
			//前面都没有左括号 就出右括号 认为不匹配
			if stack.IsEmpty() {
				result = false
				break
			} else {
				bracket := popStack(stack)
				if !isMatching(bracket, backedVal) {
					result = false
					break
				}
			}
		}
		result = true
	}

	result = result && stack.IsEmpty()
	t.Log(result)
}

func isLeftBracket(str string) bool {
	return str == "[" || str == "(" || str == "{"
}

func isRightBracket(str string) bool {
	return str == "]" || str == ")" || str == "}"
}

func isMatching(leftStr string, rightStr string) bool {
	switch leftStr {
	case "[":
		return rightStr == "]"
	case "(":
		return rightStr == ")"
	case "{":
		return rightStr == "}"
	default:
		return false
	}
}
