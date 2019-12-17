package main

import "fmt"

type MinStack struct {
	items  []int
	length int
	count  int
	min    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minsk := MinStack{}

	minsk.length = 10
	minsk.count = 0
	minsk.min = 9999999999999
	minsk.items = make([]int, minsk.length)

	return minsk
}

func (s *MinStack) Push(x int) {
	if s.count == s.length {
		return
	}
	s.items[s.count] = x
	if s.count == 0 {
		s.min = x
	} else if x <= s.min {
		s.min = x
	}
	s.count++
	return
}

func (s *MinStack) Pop() {
	if s.count == 0 {
		return
	}
	if s.count == 1 {
		s.min = 0
	} else if s.min == s.items[s.count-1] {
		s.min = s.items[0]
		for index, val := range s.items {
			if index == s.count-1 {
				break
			}
			if val < s.min {
				s.min = val
			}
		}
	}
	s.items[s.count-1] = 0
	s.count--
	return
}

func (s *MinStack) Top() int {
	if s.count == 0 {
		return 0
	}
	element := s.items[s.count-1]
	return element
}

func (s *MinStack) GetMin() int {
	return s.min
}

func main() {
	obj := Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	obj.Push(2147483647)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Push(-2147483648)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
