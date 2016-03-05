package templates

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type orderedAvlSetSuite struct{}

var _ = Suite(&orderedAvlSetSuite{})

func (s *orderedAvlSetSuite) TestContains(c *C) {
	set := NewOrderedSet()
	c.Assert(set.Contains(42), Equals, false)
	set.Add(42)
	c.Assert(set.Contains(42), Equals, true)
}

func (s *orderedAvlSetSuite) TestAddFirst(c *C) {
	set := NewOrderedSet()
	c.Assert(set.Add(42), Equals, true)
}

func (s *orderedAvlSetSuite) TestAddMany(c *C) {
	set := NewOrderedSet()
	set.Add(42)
	c.Assert(set.Add(32), Equals, true)
	c.Assert(set.Add(55), Equals, true)
	c.Assert(set.Add(42), Equals, false)
	c.Assert(set.Contains(32), Equals, true)
	c.Assert(set.Contains(55), Equals, true)
	c.Assert(set.Contains(42), Equals, true)
}
