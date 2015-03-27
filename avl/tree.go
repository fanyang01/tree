package avl

import "errors"

// Interface need to implement Compare method
type Interface interface {
	// Compare data of reciever with argument
	// Return 0 if equal
	// Return 1/-1 if greater/less
	Compare(Interface) int
}

type node struct {
	left, right *node
	data        Interface
	height      int
}

// Tree is the AVL tree
type Tree struct {
	root *node
	size int
}

// New make a new Tree
func New() *Tree {
	return &Tree{
		size: 0,
		root: nil,
	}
}

// Insert insert data into tree
func (t *Tree) Insert(data Interface) {
	n := newNode(data)
	t.root = insert(t.root, n)
	t.size++
}

// Delete remove node with given data from tree
func (t *Tree) Delete(data Interface) (Interface, error) {
	if t.root == nil {
		return nil, errors.New("delete on an empty tree")
	}
	var n *node
	t.root, n = deletion(t.root, data)
	if n == nil {
		return nil, errors.New("not found data to be deleted")
	}
	t.size--
	return n.data, nil
}

// Search return data equal to given data according to Compare method
func (t *Tree) Search(data Interface) (Interface, error) {
	if t.root == nil {
		return nil, errors.New("tree is empty")
	}
	n := search(t.root, data)
	if n == nil {
		return nil, errors.New("not found")
	}
	return n.data, nil
}

// return a new-allocated node
func newNode(data Interface) *node {
	return &node{
		data:   data,
		left:   nil,
		right:  nil,
		height: 0,
	}
}

func height(n *node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *node) balanceFactor() int {
	return height(n.left) - height(n.right)
}

func (n *node) updateHeight() {
	n.height = max(height(n.left), height(n.right)) + 1
}

func balance(r *node) *node {
	f := r.balanceFactor()
	if f == 2 {
		if r.left.balanceFactor() >= 0 {
			r = rightRotate(r)
		} else {
			r = doubleRightRotate(r)
		}
	} else if f == -2 {
		if r.right.balanceFactor() <= 0 {
			r = leftRotate(r)
		} else {
			r = doubleLeftRotate(r)
		}
	}
	return r
}

func insert(r, n *node) *node {
	if r == nil {
		return n
	}
	if n.data.Compare(r.data) < 0 {
		r.left = insert(r.left, n)
	} else {
		r.right = insert(r.right, n)
	}
	r.updateHeight()
	return balance(r)
}

func deletion(r *node, data Interface) (*node, *node) {
	if r == nil {
		return nil, nil
	}
	var cmp int
	var deleted *node
	var y *node
	if cmp = data.Compare(r.data); cmp < 0 {
		r.left, deleted = deletion(r.left, data)
	} else if cmp > 0 {
		r.right, deleted = deletion(r.right, data)
	} else {
		deleted = r
		if r.right == nil {
			return r.left, deleted
		} else if r.left == nil {
			return r.right, deleted
		} else {
			r.right, y = deleteMin(r.right)
			y.left, y.right = r.left, r.right
			r = y
		}
	}
	r.updateHeight()
	return balance(r), deleted
}

func deleteMin(r *node) (*node, *node) {
	if r.left == nil {
		return r.right, r
	}
	var deleted *node
	r.left, deleted = deleteMin(r.left)
	r.updateHeight()
	return balance(r), deleted
}

func search(r *node, data Interface) *node {
	var cmp int
	for r != nil {
		if cmp = data.Compare(r.data); cmp == 0 {
			return r
		} else if cmp < 0 {
			r = r.left
		} else {
			r = r.right
		}
	}
	return nil
}

/*
 *           x
 *          / \
 *         a   y
 *            / \
 *           b   c
 * ->
 *           y
 *          / \
 *         x   c
 *        / \
 *       a   b
 */
func leftRotate(x *node) *node {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

/*
 *           x
 *          / \
 *         y   c
 *        / \
 *       a   b
 * ->
 *           y
 *          / \
 *         a   x
 *            / \
 *           b   c
 */
func rightRotate(x *node) *node {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

/*
 *              x
 *             / \
 *            y   d
 *           / \
 *          a   z
 *             / \
 *            b   c
 *  ->
 *              z
 *            /   \
 *           y     x
 *          / \   / \
 *         a   b c   d
 */
func doubleRightRotate(x *node) *node {
	y := x.left
	z := y.right
	x.left = z.right
	y.right = z.left
	z.left = y
	z.right = x
	return z
}

/*
 *            x
 *           / \
 *          a   y
 *             / \
 *            z   d
 *           / \
 *          b   c
 * ->
 *            z
 *          /   \
 *         x     y
 *        / \   / \
 *       a   b c   d
 */
func doubleLeftRotate(x *node) *node {
	y := x.right
	z := y.left
	x.right = z.left
	y.left = z.right
	z.right = y
	z.left = x
	return z
}
