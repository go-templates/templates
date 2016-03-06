package templates

type treeNodeStack struct {
	array []*treeNode
}

func NewStack() treeNodeStack {
	s := treeNodeStack{}
	s.array = []*treeNode{}
	return treeNodeStack{}
}

func (s *treeNodeStack) Push(val *treeNode) {
	s.array = append(s.array, val)
}

func (s *treeNodeStack) Pop() (*treeNode, bool) {
	if len(s.array) == 0 {
		return nil, false
	}
	val := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return val, true
}

func (s *treeNodeStack) Peek() (*treeNode, bool) {
	if len(s.array) == 0 {
		return nil, false
	}
	val := s.array[len(s.array)-1]
	return val, true
}

func (s *treeNodeStack) IsEmpty() bool {
	if len(s.array) == 0 {
		return true
	}
	return false
}
