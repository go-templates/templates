package stack

type Stack struct {
	array []int
}

func NewStack() Stack {
	s := Stack{}
	s.array = []int{}
	return Stack{}
}

func (s *Stack) Push(val int) {
	s.array = append(s.array, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.array) == 0 {
		return 0, false
	}
	val := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.array) == 0 {
		return 0, false
	}
	val := s.array[len(s.array)-1]
	return val, true
}

func (s *Stack) IsEmpty() bool {
	if len(s.array) == 0 {
		return true
	}
	return false
}
