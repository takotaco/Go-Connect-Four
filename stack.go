package main

type StackErr struct{}

func (e StackErr) Error() string {
	return "stack is already full"
}

type Stack struct {
	items []string
	count int
	max   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v string) (ok bool, err error) {
	if s.count == s.max {
		return false, StackErr(*new(StackErr))
	}
	s.items = append(s.items[:s.count], v)
	s.count++
	return true, nil
}

func (s *Stack) Pop() (v string) {
	if s.count == 0 {
		return ""
	}
	s.count--
	v = s.items[s.count]
	s.items = s.items[:s.count]
	return v
}
