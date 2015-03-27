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
	TD234 = 0
	BU23  = 1
)

type node struct {
	left, right *node
	color       bool
	data        Interface
}

// Tree is left learning red black tree
type Tree struct {
	mode int
	size int
	root *node
}

// New return a new-allocated llrb
// mode can be TD234 or BU23, otherwise it default to be BU23
func New(mode int) *Tree {
	if mode != TD234 && mode != BU23 {
		mode = BU23
	}
	return &Tree{
		size: 0,
		root: nil,
		mode: mode,
	}
}

// Insert insert data into correct place
func (t *Tree) Insert(data Interface) {
	n := newNode(data)
	t.insert(&t.root, n)
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
func (t *Tree) insert(pr **node, n *node) {
	r := *pr
	if r == nil {
		*pr = n
		return
	}
	// TD234, split 4-node into 2-2 nodes
	if t.mode == TD234 && isRed(r.left) && isRed(r.right) {
		colorFlip(r)
	}
	var cmp int
	if cmp = n.data.Compare(r.data); cmp < 0 {
		t.insert(&r.left, n)
	} else {
		t.insert(&r.right, n)
	}

	// left leaning
	if isRed(r.right) {
		leftRotate(pr)
		r = *pr
	}
	// successive left leaning
	if isRed(r.left) && isRed(r.left.left) {
		rightRotate(pr)
		r = *pr
	}
	// BU23, split 4-node into 2-2 nodes
	if t.mode == BU23 && isRed(r.left) && isRed(r.right) {
		colorFlip(r)
	}
}

func isRed(n *node) bool {
	if n == nil {
		return false
	}
	return n.color == RED
}

func colorFlip(n *node) {
	n.color = !n.color
	if n.left != nil {
		n.left.color = !n.left.color
	}
	if n.left != nil {
		n.right.color = !n.right.color
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
