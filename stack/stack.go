package stack

import "errors"

type Stack struct {
	Top   int
	Items []int
	Size  int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data int) {
	s.Top = data
	s.Items = append(s.Items, data)
	s.Size++
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	topIndex := s.Size - 1
	value := s.Items[topIndex]

	s.Items = s.Items[:topIndex]
	s.Size--

	if s.Size > 0 {
		s.Top = s.Items[s.Size-1]
	} else {
		s.Top = 0
	}

	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.Top, nil
}
