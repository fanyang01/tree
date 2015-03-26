package llrb

import "errors"

// Delete remove and return given data from llrb, but return error if data not found
func (t *Tree) Delete(data Interface) (Interface, error) {
	n := deletion(&t.root, data)
	if n == nil {
		return nil, errors.New("not found")
	}
	t.root.color = BLACK
	t.size--
	return n.data, nil
}

// delete conflict with delete in go
func deletion(pr **node, data Interface) *node {
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
		ret = deletion(&r.left, data)
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
			y := deleteMin(&r.right)
			y.left, y.right = r.left, r.right
			y.color = r.color
			*pr = y
		} else {
			ret = deletion(&r.right, data)
		}
	}
	fixup(pr)
	return ret
}

func moveRedLeft(pr **node) {
	r := *pr
	colorFlip(r)
	if isRed(r.right.left) {
		rightRotate(&r.right)
		leftRotate(pr)
		colorFlip(*pr)
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

func deleteMin(pr **node) *node {
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
	ret = deleteMin(&r.left)
	fixup(pr)
	return ret
}
