package stackList

import "errors"

type StackList struct {
	items  []string
	count  int
	length int
}

func (s *StackList) CreateStack(n int) {
	s.length = n
	s.items = make([]string, n)
	s.count = 0
}

func (s *StackList) Push(element string) (bool, error) {
	if s.count == s.length {
		return false, errors.New("stack full")
	}
	s.items[s.count] = element
	s.count++
	return true, nil
}

func (s *StackList) Pop() (string, error) {
	if s.count == 0 {
		return "", errors.New("stack empty")
	}
	element := s.items[s.count-1]
	s.items[s.count-1] = ""
	s.count--
	return element, nil
}

func (s *StackList) Top() (string, error) {
	if s.count == 0 {
		return "", errors.New("stack empty")
	}
	element := s.items[s.count-1]
	return element, nil
}

func (s *StackList) IsEmpty() bool {
	return s.count == 0
}
