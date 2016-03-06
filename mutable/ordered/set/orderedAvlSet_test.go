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

func (s *orderedAvlSetSuite) TestSetUniqueness(c *C) {
	set := NewOrderedSet()
	set.Add(42)
	c.Assert(set.Add(32), Equals, true)
	c.Assert(set.Add(55), Equals, true)
	c.Assert(set.Add(42), Equals, false)
	c.Assert(set.Contains(32), Equals, true)
	c.Assert(set.Contains(55), Equals, true)
	c.Assert(set.Contains(42), Equals, true)
}

func (s *orderedAvlSetSuite) TestDoubleRotateRight(c *C) {
	set := NewOrderedSet()
	set.Add(9)
	set.Add(5)
	set.Add(8)
	c.Assert(set.root.height, Equals, 2)
	c.Assert(set.root.left.height, Equals, 1)
	c.Assert(set.root.value, Equals, 8)
	c.Assert(set.root.right.height, Equals, 1)
}

func (s *orderedAvlSetSuite) TestRotateRight(c *C) {
	set := NewOrderedSet()
	set.Add(9)
	set.Add(8)
	set.Add(5)
	c.Assert(set.root.height, Equals, 2)
	c.Assert(set.root.value, Equals, 8)
	c.Assert(set.root.left.height, Equals, 1)
	c.Assert(set.root.right.height, Equals, 1)
}

func (s *orderedAvlSetSuite) TestDoubleRotateLeft(c *C) {
	set := NewOrderedSet()
	set.Add(5)
	set.Add(9)
	set.Add(8)
	c.Assert(set.root.value, Equals, 8)
	c.Assert(set.root.height, Equals, 2)
	c.Assert(set.root.left.height, Equals, 1)
	c.Assert(set.root.right.height, Equals, 1)
}

func (s *orderedAvlSetSuite) TestDoubleRotateLeftLower(c *C) {
	set := NewOrderedSet()
	set.Add(12)
	set.Add(9)
	set.Add(15)
	set.Add(13)
	set.Add(14)
	c.Assert(set.root.value, Equals, 12)
	c.Assert(set.root.left.value, Equals, 9)
	c.Assert(set.root.right.left.value, Equals, 13)
	c.Assert(set.root.right.right.value, Equals, 15)
	c.Assert(set.root.right.value, Equals, 14)
	c.Assert(set.root.height, Equals, 3)
	c.Assert(set.root.left.height, Equals, 1)
	c.Assert(set.root.right.height, Equals, 2)
}

func (s *orderedAvlSetSuite) TestSize(c *C) {
	set := NewOrderedSet()
	c.Assert(set.Size(), Equals, 0)
	set.Add(5)
	c.Assert(set.Size(), Equals, 1)
	set.Add(5)
	c.Assert(set.Size(), Equals, 1)
	set.Add(9)
	c.Assert(set.Size(), Equals, 2)
}

func (s *orderedAvlSetSuite) TestRemoveNonExisting(c *C) {
	set := NewOrderedSet()
	c.Assert(set.Remove(5), Equals, false)
}
