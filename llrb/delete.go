package llrb

import "errors"

// Delete remove and return given data from llrb, but return error if data not found
func (t *Tree) Delete(data Interface) (Interface, error) {
	n := t.deletion(&t.root, data)
	if n == nil {
		return nil, errors.New("not found")
	}
	if t.root != nil {
		t.root.color = BLACK
	}
	t.size--
	return n.data, nil
}

// delete conflict with delete in go
func (t *Tree) deletion(pr **node, data Interface) *node {
	var ret *node

	r := *pr
	if r == nil {
		return nil
	}
	cmp := data.Compare(r.data)
	if cmp < 0 {
		if r.left == nil {
			return nil
		}
		if !isRed(r.left) && !isRed(r.left.left) {
			moveRedLeft(pr)
			r = *pr
		}
		ret = t.deletion(&r.left, data)
	} else {
		if isRed(r.left) {
			rightRotate(pr)
			r = *pr
			cmp = data.Compare(r.data)
		}
		if cmp == 0 && r.right == nil {
			*pr = nil
			return r
		} else if r.right == nil {
			return nil
		}
		if !isRed(r.right) && !isRed(r.right.left) {
			moveRedRight(pr)
			r = *pr
			cmp = data.Compare(r.data)
		}
		if cmp == 0 {
			ret = r
			y := t.deleteMin(&r.right)
			y.left, y.right = r.left, r.right
			y.color = r.color
			*pr = y
		} else {
			ret = t.deletion(&r.right, data)
		}
	}
	fixup(pr, t.mode)
	return ret
}

func fixup(pr **node, mode int) {
	r := *pr
	// left leaning
	if isRed(r.right) {
		if mode == TD234 && isRed(r.right.left) {
			rightRotate(&r.right)
		}
		leftRotate(pr)
		r = *pr
	}
	// successive left leaning
	if isRed(r.left) && isRed(r.left.left) {
		rightRotate(pr)
		r = *pr
	}
	// BU23, split 4-node into 2-2 nodes
	if mode == BU23 && isRed(r.left) && isRed(r.right) {
		colorFlip(r)
	}
}

func moveRedLeft(pr **node) {
	r := *pr
	colorFlip(r)
	if isRed(r.right.left) {
		rightRotate(&r.right)
		leftRotate(pr)
		colorFlip(*pr)
		// for TD234
		r = *pr
		if r.right != nil && isRed(r.right.right) {
			leftRotate(&r.right)
		}
	}
}

func moveRedRight(pr **node) {
	r := *pr
	colorFlip(r)
	if isRed(r.left.left) {
		rightRotate(pr)
		colorFlip(*pr)
	}
}

func (t *Tree) deleteMin(pr **node) *node {
	var ret *node
	r := *pr
	if r.left == nil {
		// somthing error
		*pr = nil
		return r
	}
	if !isRed(r.left) && !isRed(r.left.left) {
		moveRedLeft(pr)
		r = *pr
	}
	ret = t.deleteMin(&r.left)
	fixup(pr, t.mode)
	return ret
}
