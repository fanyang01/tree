package splay

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
}

// flags used by splay
const (
	LEFT       = 0
	RIGHT      = 1
	LEFTLEFT   = LEFT<<1 | LEFT
	LEFTRIGHT  = LEFT<<1 | RIGHT
	RIGHTLEFT  = RIGHT<<1 | LEFT
	RIGHTRIGHT = RIGHT<<1 | RIGHT
)

// Tree is the splay tree
type Tree struct {
	size int
	root *node
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
	if t.root == nil {
		t.root = n
	} else {
		t.root = zigRoot(insert(t.root, n))
	}
	t.size++
}

// Search return found data if successful, or return error when tree is empty or data not found
func (t *Tree) Search(data Interface) (Interface, error) {
	if t.size <= 0 {
		return nil, errors.New("emtpy tree")
	}
	ret, _ := t.Splay(data)
	if t.root.data.Compare(data) != 0 {
		return nil, errors.New("not found")
	}
	return ret, nil
}

// Delete remove node with given data from tree
func (t *Tree) Delete(data Interface) (Interface, error) {
	r, path, d := deletion(t.root, data)
	if d == nil {
		return nil, errors.New("can't find data to be deleted")
	}
	t.root = zigRoot(r, path)
	t.size--
	return d, nil
}

func insert(r, n *node) (*node, int) {
	if r == nil {
		return n, -1
	}
	var (
		cmp, path, prev int
	)
	if cmp = n.data.Compare(r.data); cmp < 0 {
		r.left, prev = insert(r.left, n)
		path = LEFT
	} else {
		r.right, prev = insert(r.right, n)
		path = RIGHT
	}
	if prev == -1 {
		return r, path
	}
	path = prev<<1 | path
	return splayPath(r, path)
}

func deletion(r *node, data Interface) (*node, int, Interface) {
	if r == nil {
		return nil, -2, nil
	}
	var (
		cmp, path, prev int
		ret             Interface
	)
	if cmp = data.Compare(r.data); cmp < 0 {
		r.left, prev, ret = deletion(r.left, data)
		path = LEFT
	} else if cmp > 0 {
		r.right, prev, ret = deletion(r.right, data)
		path = RIGHT
	} else {
		if r.left == nil {
			return r.right, -2, r.data
		}
		if r.right == nil {
			return r.left, -2, r.data
		}
		var y *node
		r.right, y = deleteMin(r.right)
		y.left, y.right = r.left, r.right
		return y, -2, r.data
	}
	if prev == -2 {
		return r, -1, ret
	}
	if prev == -1 {
		return r, path, ret
	}
	path = prev<<1 | path
	r, path = splayPath(r, path)
	return r, path, ret

}

func deleteMin(r *node) (*node, *node) {
	r = zigRoot(splay(r, nil, -1))
	return r.right, r
}

// Splay splay the node with data to root
// similiar as Search, but when search a non-exist node, the return data is not valid
func (t *Tree) Splay(data Interface) (Interface, error) {
	if t.root == nil {
		return nil, errors.New("empty tree")
	}
	t.root = zigRoot(splay(t.root, data, 0))
	return t.root.data, nil
}

// Join two trees togather
func Join(a, b *Tree) (*Tree, error) {
	if a == nil && b == nil {
		return nil, errors.New("can't join two nil Trees")
	}
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}
	a.root = join(a.root, b.root)
	a.size += b.size
	return a, nil
}

func join(l, r *node) *node {
	if l == nil {
		return r
	}
	// splay the maximum node of l to root
	// after this the root of l have no right childs
	l = zigRoot(splay(l, nil, 1))
	l.right = r
	return l
}

// when flag is 0, compare a and b
// when flag is not 0, return flag
func compare(a, b Interface, flag int) int {
	if flag == 0 {
		return a.Compare(b)
	}
	return flag
}

// return a new-allocated node
func newNode(data Interface) *node {
	return &node{
		data:  data,
		left:  nil,
		right: nil,
	}
}
