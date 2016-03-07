package templates

type OrderedSet interface {
	Add(val int) bool
	Remove(val int) bool
	Contains(val int) bool
	Size() int
	Iter() OrderedSetIterator
}

type OrderedSetIterator interface {
	Value() int
	Next() bool
}

type orderedAvlSet struct {
	root *treeNode
	size int
}

type treeNode struct {
	left   *treeNode
	right  *treeNode
	value  int
	height int
}

type orderedAvlSetIterator struct {
	history     treeNodeStack
	currentNode *treeNode
}

func (o *orderedAvlSet) Iter() orderedAvlSetIterator {
	iter := orderedAvlSetIterator{}
	iter.currentNode = o.root
	if iter.currentNode != nil {
		for iter.currentNode.left != nil {
			iter.history.Push(iter.currentNode)
			iter.currentNode = iter.currentNode.left
		}
	}
	return iter
}

func (iter *orderedAvlSetIterator) Value() int {
	if iter.currentNode == nil {
		return 0
	}
	return iter.currentNode.value
}

func (iter *orderedAvlSetIterator) Next() bool {
	if iter.currentNode != nil {
		if iter.currentNode.right != nil {
			iter.currentNode = iter.currentNode.right
			for iter.currentNode.left != nil {
				iter.currentNode = iter.currentNode.left
			}
		} else {
			iter.currentNode, _ = iter.history.Pop()
			if iter.currentNode == nil {
				return false
			}
		}
		return true
	}
	return false
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
	if o.root == nil {
		o.root = &treeNode{left: nil, right: nil, value: val, height: 1}
		o.size = 1
		return true
	}
	added := false
	o.root, added = o.root.add(val)
	if added {
		o.size++
	}
	return added
}

func (currentNode *treeNode) add(val int) (*treeNode, bool) {
	added := false
	if currentNode.value < val {
		if currentNode.right == nil {
			currentNode.right = &treeNode{left: nil, right: nil, value: val, height: 1}
			added = true
		} else {
			currentNode.right, added = currentNode.right.add(val)
		}
	} else if currentNode.value > val {
		if currentNode.left == nil {
			currentNode.left = &treeNode{left: nil, right: nil, value: val, height: 1}
			added = true
		} else {
			currentNode.left, added = currentNode.left.add(val)
		}
	}
	currentNode = currentNode.balance()
	return currentNode, added
}

func (o *orderedAvlSet) Remove(val int) bool {
	removed := false
	if o.root == nil {
		return removed
	} else {
		o.root, removed = o.root.remove(val)
	}
	if removed {
		o.size--
	}
	return removed
}

func (currentNode *treeNode) remove(val int) (*treeNode, bool) {
	removed := false
	if currentNode.value == val {
		if currentNode.left != nil && currentNode.right != nil {
			currentNode.value = findMax(currentNode.left)
			currentNode.left, removed = currentNode.left.remove(currentNode.value)
		} else if currentNode.left != nil {
			currentNode = currentNode.left
			removed = true
		} else if currentNode.right != nil {
			currentNode = currentNode.right
			removed = true
		} else if currentNode.right == nil && currentNode.left == nil {
			currentNode = nil
			removed = true
			return currentNode, removed
		}
	} else if currentNode.value < val {
		currentNode.right, removed = currentNode.right.remove(val)
	} else if currentNode.value > val {
		currentNode.left, removed = currentNode.left.remove(val)
	}
	currentNode = currentNode.balance()
	return currentNode, removed
}

func findMax(currentNode *treeNode) int {
	if currentNode.right != nil {
		return findMax(currentNode.right)
	}
	return currentNode.value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getHeight(node *treeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}
func (currentNode *treeNode) balance() *treeNode {
	if getHeight(currentNode.right)-getHeight(currentNode.left) == 2 {
		if getHeight(currentNode.right.right) >= getHeight(currentNode.right.left) {
			currentNode = currentNode.singleRotateLeft()
		} else {
			currentNode = currentNode.doubleRotateLeft()
		}
	} else if getHeight(currentNode.left)-getHeight(currentNode.right) == 2 {
		if getHeight(currentNode.left.left) >= getHeight(currentNode.left.right) {
			currentNode = currentNode.singleRotateRight()
		} else {
			currentNode = currentNode.doubleRotateRight()
		}
	}
	currentNode.height = max(getHeight(currentNode.left), getHeight(currentNode.right)) + 1
	return currentNode
}

func (currentNode *treeNode) singleRotateRight() *treeNode {
	newNode := currentNode.left
	currentNode.left = newNode.right
	newNode.right = currentNode
	currentNode.height = max(getHeight(currentNode.left), getHeight(currentNode.right)) + 1
	newNode.height = max(getHeight(newNode.left), currentNode.height) + 1
	return newNode
}

func (currentNode *treeNode) doubleRotateRight() *treeNode {
	currentNode.left = currentNode.left.singleRotateLeft()
	return currentNode.singleRotateRight()
}

func (currentNode *treeNode) singleRotateLeft() *treeNode {
	newNode := currentNode.right
	currentNode.right = newNode.left
	newNode.left = currentNode
	currentNode.height = max(getHeight(currentNode.left), getHeight(currentNode.right)) + 1
	newNode.height = max(getHeight(newNode.right), currentNode.height) + 1
	return newNode
}

func (currentNode *treeNode) doubleRotateLeft() *treeNode {
	currentNode.right = currentNode.right.singleRotateRight()
	return currentNode.singleRotateLeft()
}
