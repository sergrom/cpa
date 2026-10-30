package ds

// AVLTree is a balanced multiset of integers. Its zero value is ready to use.
type AVLTree struct {
	root *avlNode
	len  int
}

// NewAVLTree returns an empty AVL tree.
func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// Insert inserts a new value into the tree and returns the tree pointer
func (t *AVLTree) Insert(value int) *AVLTree {
	t.root = t.root.insert(value)
	t.len++
	return t
}

// Get returns the node value associated with the search value
func (t *AVLTree) Get(value int) (int, bool) {
	var node *avlNode
	if t.root != nil {
		node = t.root.get(value)
	}
	if node != nil {
		return node.Value, true
	}
	return 0, false
}

// Len returns the number of values, including duplicates.
func (t *AVLTree) Len() int {
	return t.len
}

// Head returns the first value in the tree
func (t *AVLTree) Head() (int, bool) {
	if t.root == nil {
		return 0, false
	}
	var beginning = t.root
	for beginning.left != nil {
		beginning = beginning.left
	}
	return beginning.Value, true
}

// Tail returns the last value in the tree
func (t *AVLTree) Tail() (int, bool) {
	if t.root == nil {
		return 0, false
	}
	var beginning = t.root
	for beginning.right != nil {
		beginning = beginning.right
	}
	return beginning.Value, true
}

// Delete removes one occurrence of value, if present.
func (t *AVLTree) Delete(value int) *AVLTree {
	deleted := false
	t.root = t.root.deleteNode(value, &deleted)
	if deleted {
		t.len--
	}
	return t
}

// Pop removes one occurrence of the maximum value.
func (t *AVLTree) Pop() (int, bool) {
	value, ok := t.Tail()
	if ok {
		t.Delete(value)
	}
	return value, ok
}

// Pull removes one occurrence of the minimum value.
func (t *AVLTree) Pull() (int, bool) {
	value, ok := t.Head()
	if ok {
		t.Delete(value)
	}
	return value, ok
}

func (t *AVLTree) balance() int8 {
	if t.root != nil {
		return t.root.balance()
	}
	return 0
}

// =============== Node ================

// Node represents a node in the tree with a value, left and right children, and a height/balance of the node.
type avlNode struct {
	count       int
	Value       int
	left, right *avlNode
	hght        int8
}

// Init initializes the values of the node or clears the node and returns the node pointer
func (n *avlNode) init() *avlNode {
	n.left, n.right = nil, nil
	n.hght = 1
	n.count = 1
	return n
}

func (n *avlNode) get(val int) *avlNode {
	var node *avlNode
	if c := cmp(val, n.Value); c < 0 {
		if n.left != nil {
			node = n.left.get(val)
		}
	} else if c > 0 {
		if n.right != nil {
			node = n.right.get(val)
		}
	} else {
		node = n
	}
	return node
}

func (n *avlNode) rotateRight() *avlNode {
	l := n.left
	// Rotation
	l.right, n.left = n, l.right

	// update heights
	n.hght = n.maxHeight() + 1
	l.hght = l.maxHeight() + 1

	return l
}

func (n *avlNode) rotateLeft() *avlNode {
	r := n.right
	// Rotation
	r.left, n.right = n, r.left

	// update heights
	n.hght = n.maxHeight() + 1
	r.hght = r.maxHeight() + 1

	return r
}

func (n *avlNode) min() *avlNode {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *avlNode) maxHeight() int8 {
	rh, lh := n.right.height(), n.left.height()
	if rh > lh {
		return rh
	}
	return lh
}

func (n *avlNode) height() int8 {
	if n != nil {
		return n.hght
	}
	return 0
}

func (n *avlNode) balance() int8 {
	if n == nil {
		return 0
	}
	return n.left.height() - n.right.height()
}

func (n *avlNode) insert(value int) *avlNode {
	if n == nil {
		return (&avlNode{Value: value}).init()
	}
	c := cmp(value, n.Value)
	if c > 0 {
		n.right = n.right.insert(value)
	} else if c < 0 {
		n.left = n.left.insert(value)
	} else {
		n.count++
		return n
	}

	n.hght = n.maxHeight() + 1

	if c = n.balance(); c > 1 {
		c = cmp(value, n.left.Value)
		if c < 0 {
			return n.rotateRight()
		} else if c > 0 {
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	} else if c < -1 {
		c = cmp(value, n.right.Value)
		if c > 0 {
			return n.rotateLeft()
		} else if c < 0 {
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}
	return n
}

func (n *avlNode) deleteNode(value int, deleted *bool) *avlNode {
	if n == nil {
		return n
	}

	if c := cmp(value, n.Value); c < 0 {
		n.left = n.left.deleteNode(value, deleted)
	} else if c > 0 {
		n.right = n.right.deleteNode(value, deleted)
	} else {
		*deleted = true
		if n.count > 1 {
			n.count--
			return n
		}
		if n.left == nil {
			t := n.right
			n.init()
			return t
		} else if n.right == nil {
			t := n.left
			n.init()
			return t
		}
		t := n.right.min()
		n.Value, n.count = t.Value, t.count
		t.count = 1 // Move every occurrence, then remove the successor node.
		n.right = n.right.deleteNode(t.Value, deleted)
	}

	// re-balance
	if n == nil {
		return n
	}
	n.hght = n.maxHeight() + 1
	bal := n.balance()
	if bal > 1 {
		if n.left.balance() >= 0 {
			return n.rotateRight()
		}
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	} else if bal < -1 {
		if n.right.balance() <= 0 {
			return n.rotateLeft()
		}
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func cmp(a, b int) int8 {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// Count returns the number of occurrences of value.
func (t *AVLTree) Count(value int) int {
	if t.root != nil {
		if n := t.root.get(value); n != nil {
			return n.count
		}
	}
	return 0
}

// Lower returns the greatest value strictly less than value.
func (t *AVLTree) Lower(value int) (int, bool) {
	var result int
	found := false
	for n := t.root; n != nil; {
		if n.Value < value {
			result, found = n.Value, true
			n = n.right
		} else {
			n = n.left
		}
	}
	return result, found
}

// Higher returns the smallest value strictly greater than value.
func (t *AVLTree) Higher(value int) (int, bool) {
	var result int
	found := false
	for n := t.root; n != nil; {
		if n.Value > value {
			result, found = n.Value, true
			n = n.left
		} else {
			n = n.right
		}
	}
	return result, found
}

// Ascend visits values in ascending order, including duplicates, until visit returns false.
// The callback must not modify the tree during traversal.
func (t *AVLTree) Ascend(visit func(int) bool) {
	var walk func(*avlNode) bool
	walk = func(n *avlNode) bool {
		if n == nil {
			return true
		}
		if !walk(n.left) {
			return false
		}
		for i := 0; i < n.count; i++ {
			if !visit(n.Value) {
				return false
			}
		}
		return walk(n.right)
	}
	walk(t.root)
}

// Values returns an independent sorted slice containing all values, including duplicates.
func (t *AVLTree) Values() []int {
	values := make([]int, 0, t.len)
	t.Ascend(func(value int) bool { values = append(values, value); return true })
	return values
}

// tree := ds.NewAVLTree()
// tree.Insert(3).Insert(1).Insert(3).Insert(7)

// tree.Values()  // [1 3 3 7]
// tree.Len()     // 4
// tree.Count(3)  // 2
// tree.Lower(3)  // 1, true — строго меньше
// tree.Higher(3) // 7, true — строго больше

// tree.Delete(3) // удаляет одно вхождение

// tree.Ascend(func(value int) bool {
//     fmt.Println(value)
//     return true // false остановит обход
// })
