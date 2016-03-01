package forward

type ForwardList interface {
	Prepend(val int) ForwardList
	First() int
	Iter() ForwardListIterator
}

type ForwardListIterator interface {
	Value() int
	Next() bool
}

type forwardList struct {
	next  *forwardList
	value int
}

type forwardListIterator struct {
	cur *forwardList
	idx int
}

func NewForwardList() *forwardList {
	return nil
}

func (f *forwardList) Prepend(val int) *forwardList {
	return &forwardList{value: val, next: f}
}

func (f *forwardList) Iter() *forwardListIterator {
	return &forwardListIterator{f, 0}
}

func (f *forwardListIterator) Next() bool {
	if f.cur.next == nil {
		return false
	}
	if f.idx == 0 {
		f.idx++
	} else {
		f.cur = f.cur.next
	}
	return true
}

func (f *forwardListIterator) Value() int {
	if f.cur == nil {
		return 0
	}
	return f.cur.value
}
