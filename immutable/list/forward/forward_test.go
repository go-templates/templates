package forward

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ForwardSuite struct{}

var _ = Suite(&ForwardSuite{})

func (s *ForwardSuite) TestAdd(c *C) {
	l := NewForwardList()
	l2 := l.Prepend(23)
	c.Assert(l2.Iter().Value(), Equals, 23)
}

func (s *ForwardSuite) TestRange(c *C) {
	l := NewForwardList()
	l2 := l.Prepend(23)
	l3 := l2.Prepend(42)
	l4 := l3.Prepend(47)
	iter := l4.Iter()
	for i := 0; iter.Next(); i++ {
		switch i {
		case 0:
			c.Assert(iter.Value(), Equals, 47)
		case 1:
			c.Assert(iter.Value(), Equals, 42)
		case 2:
			c.Assert(iter.Value(), Equals, 23)
		}
	}
}

func (s *ForwardSuite) TestEmptyIter(c *C) {
	l := NewForwardList()
	c.Assert(l.Iter().Value(), Equals, 0)
	c.Assert(l.Iter().Next(), Equals, false)
}

func (s *ForwardSuite) TestGet(c *C) {
	l := NewForwardList()
	l2 := l.Prepend(23)
	l3 := l2.Prepend(42)
	l4 := l3.Prepend(47)
	el, ok := l4.Get(2)
	c.Assert(ok, Equals, true)
	c.Assert(el, Equals, 23)
	el, ok = l4.Get(0)
	c.Assert(ok, Equals, true)
	c.Assert(el, Equals, 47)
	el, ok = l4.Get(1)
	c.Assert(ok, Equals, true)
	c.Assert(el, Equals, 42)
}

func (s *ForwardSuite) BenchmarkPrepend(c *C) {
	l := NewForwardList()
	for i := 0; i < c.N; i++ {
		l = l.Prepend(i)
	}
}

func (s *ForwardSuite) BenchmarkGetLast(c *C) {
	l := NewForwardList()
	i := 0
	// build list
	for ; i < c.N; i++ {
		l = l.Prepend(i)
	}
	c.ResetTimer()
	// get last
	for j := 0; j < c.N; j++ {
		l.Get(i)
	}
}
