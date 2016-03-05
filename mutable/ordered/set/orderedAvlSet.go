package templates

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
	left   *treeNode
	right  *treeNode
	value  int
	height int
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
			_, added = currentNode.right.add(val)
			if currentNode.right.height-getHeight(currentNode.left) == 2 {
				if currentNode.right.value < val {
					currentNode = currentNode.singleRotateLeft()
				} else {
					currentNode = currentNode.doubleRotateLeft()
				}
			}
		}
	} else if currentNode.value > val {
		if currentNode.left == nil {
			currentNode.left = &treeNode{left: nil, right: nil, value: val, height: 1}
			added = true
		} else {
			_, added = currentNode.left.add(val)
			if currentNode.left.height-getHeight(currentNode.right) == 2 {
				if currentNode.left.value > val {
					currentNode = currentNode.singleRotateRight()
				} else {
					currentNode = currentNode.doubleRotateRight()
				}
			}
		}
	}
	currentNode.height = max(getHeight(currentNode.left), getHeight(currentNode.right)) + 1
	return currentNode, added
}
func (o *orderedAvlSet) Remove(val int) bool {
	return false
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
func (currentNode *treeNode) singleRotateRight() *treeNode {
	newNode := currentNode.left
	currentNode.left = nil
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
	currentNode.right = nil
	newNode.left = currentNode
	currentNode.height = max(getHeight(currentNode.left), getHeight(currentNode.right)) + 1
	newNode.height = max(getHeight(newNode.right), currentNode.height) + 1
	return newNode
}
func (currentNode *treeNode) doubleRotateLeft() *treeNode {
	currentNode.right = currentNode.right.singleRotateRight()
	return currentNode.singleRotateLeft()
}
