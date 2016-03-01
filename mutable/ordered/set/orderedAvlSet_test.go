package templates

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type orderedAvlSetSuite struct{}

var _ = Suite(&orderedAvlSetSuite{})

func (s *orderedAvlSetSuite) TestContains(c *C) {
	set := NewOrderedSet()
	c.Assert(set.Contains(42), Equals, false)
}

func (s *orderedAvlSetSuite) TestAddFirst(c *C) {
	set := NewOrderedSet()
	c.Assert(set.Add(42), Equals, true)
}
