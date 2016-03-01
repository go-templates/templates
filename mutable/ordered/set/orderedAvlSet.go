//ordered sets
//avl tree
//int

package templates

import ()

type OrderedSet interface {
	Add(val int) bool
	Remove(val int) bool
	Contains(val int) bool
	Size() int
}

type orderedAvlSet struct {
	root *treeNode
	size int
}

type treeNode struct {
	left  *treeNode
	right *treeNode
	value int
}

func NewOrderedSet() *orderedAvlSet {
	return &orderedAvlSet{}
}

func (o *orderedAvlSet) Size() int {
	return o.size
}

func (o *orderedAvlSet) Contains(val int) bool {
	currentNode := o.root
	for currentNode != nil {
		if currentNode.value < val {
			currentNode = currentNode.right
		} else if currentNode.value > val {
			currentNode = currentNode.left
		} else {
			return true
		}
	}
	return false
}

func (o *orderedAvlSet) Add(val int) bool {
	currentNode := o.root
	if currentNode == nil {
		o.root = &treeNode{left: nil, right: nil, value: val}
		return true
	}
	return false
}

func (o *orderedAvlSet) Remove(val int) bool {
	return false
}
