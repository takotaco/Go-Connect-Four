package stack

type StackErr struct{}

func (e StackErr) Error() string {
	return "Column is already full"
}

type Stack struct {
	items []string
	count int
	max   int
}

func New(max int) Stack {
	return Stack{
		max: max,
	}
}

func (s *Stack) Push(v string) (ok bool, err error) {
	if s.count == s.max {
		return false, new(StackErr)
	}
	s.items = append(s.items[:s.count], v)
	s.count++
	return true, nil
}

func (s Stack) IsFull() bool {
	return s.count == s.max
}

func (s Stack) Count() int {
	return s.count
}

func (s Stack) GetValueAt(i int) string {
	if i >= s.count || i < 0 {
		return ""
	} else {
		return s.items[i]
	}
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
