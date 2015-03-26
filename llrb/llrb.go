package llrb

import "errors"

// Interface need to implement Compare method
type Interface interface {
	// Compare data of reciever with argument
	// Return 0 if equal
	// Return 1/-1 if greater/less
	Compare(Interface) int
}

// color of node
const (
	BLACK = false
	RED   = true
)

type node struct {
	left, right *node
	color       bool
	data        Interface
}

// Tree is left learning red black tree
type Tree struct {
	size int
	root *node
}

// New return a new-allocated llrb
func New() *Tree {
	return &Tree{
		size: 0,
		root: nil,
	}
}

// Insert insert data into correct place
func (t *Tree) Insert(data Interface) {
	n := newNode(data)
	insert(&t.root, n)
	t.root.color = BLACK
	t.size++
}

// Search return data if found, otherwise return nil and 'not found' error
func (t *Tree) Search(data Interface) (Interface, error) {
	pr := search(&t.root, data)
	if pr == nil {
		return nil, errors.New("not found")
	}
	return (*pr).data, nil
}

func search(pr **node, data Interface) **node {
	r := *pr
	var cmp int

	for {
		if r == nil {
			break
		}
		if cmp = data.Compare(r.data); cmp == 0 {
			return pr
		}
		if cmp < 0 {
			pr = &r.left
			r = r.left
		} else {
			pr = &r.right
			r = r.right
		}
	}
	return nil
}

// pr is the address of pointer to root of subtree
// n is the new node to be inserted
func insert(pr **node, n *node) {
	r := *pr
	if r == nil {
		*pr = n
		return
	}
	var cmp int
	if cmp = n.data.Compare(r.data); cmp < 0 {
		insert(&r.left, n)
	} else {
		insert(&r.right, n)
	}
	fixup(pr)
}

func isRed(n *node) bool {
	if n == nil {
		return false
	}
	return n.color == RED
}

func colorFlip(n *node) {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

func fixup(pr **node) {
	r := *pr
	if isRed(r.right) {
		leftRotate(pr)
		r = *pr
	}
	if isRed(r.left) && isRed(r.left.left) {
		rightRotate(pr)
		r = *pr
	}
	if isRed(r.left) && isRed(r.right) {
		colorFlip(r)
	}
}

// return a new-allocated node
func newNode(data Interface) *node {
	return &node{
		data:  data,
		left:  nil,
		right: nil,
		color: RED,
	}
}

// In LLRB-Tree's rotation, we also need to maintain the color of node
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
func leftRotate(pr **node) {
	x := *pr
	y := x.right
	x.right = y.left
	y.left = x
	y.color = x.color
	x.color = RED
	*pr = y
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
func rightRotate(pr **node) {
	x := *pr
	y := x.left
	x.left = y.right
	y.right = x
	y.color = x.color
	x.color = RED
	*pr = y
}
