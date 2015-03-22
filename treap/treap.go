package treap

import (
	"errors"
	"math/rand"
)

// Interface need to implement Compare method
type Interface interface {
	// Compare data of reciever with argument
	// Return 0 if equal
	// Return 1/-1 if greater/less
	Compare(Interface) int
}

type node struct {
	left, right *node
	priority    int
	data        Interface
}

// Treap is tree-heap
type Treap struct {
	size int
	root *node
}

// New return a new-allocated treap
func New() *Treap {
	return &Treap{
		size: 0,
		root: nil,
	}
}

// Insert insert data into correct place
func (t *Treap) Insert(data Interface) {
	n := newNode(data)
	insert(&t.root, n)
	t.size++
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
		if r.priority < r.left.priority {
			rightRotate(pr)
		}
	} else {
		insert(&r.right, n)
		if r.priority < r.right.priority {
			leftRotate(pr)
		}
	}
}

// Search return data if found, otherwise return nil and 'not found' error
func (t *Treap) Search(data Interface) (Interface, error) {
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

// Delete remove and return given data from treap, but return error if data not found
func (t *Treap) Delete(data Interface) (Interface, error) {
	pr := search(&t.root, data)
	if pr == nil {
		return nil, errors.New("not found")
	}

	x := *pr
	if x.left == nil {
		*pr = x.right
	} else if x.right == nil {
		*pr = x.left
	} else {
		py := minimum(&x.right)
		y := *py
		if x.right != y {
			*py = y.right
			y.right = x.right
		}
		y.left = x.left
		y.priority = x.priority
		*pr = y
	}
	t.size--
	return x.data, nil
}

// find nimimum node in subtree
func minimum(pr **node) **node {
	r := *pr
	for {
		if r.left == nil {
			return pr
		}
		pr = &r.left
		r = r.left
	}
}

// return a new-allocated node
func newNode(data Interface) *node {
	return &node{
		data:     data,
		left:     nil,
		right:    nil,
		priority: rand.Int(),
	}
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
func leftRotate(pr **node) {
	x := *pr
	y := x.right
	x.right = y.left
	y.left = x
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
	y.right = x.left
	*pr = y
}
