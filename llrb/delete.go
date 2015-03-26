package llrb

import "errors"

// Delete remove and return given data from llrb, but return error if data not found
func (t *Tree) Delete(data Interface) (Interface, error) {
	if pr := search(&t.root, data); pr == nil {
		return nil, errors.New("not found")
	}
	ret := deletion(&t.root, data)
	t.root.color = BLACK
	t.size--
	return ret, nil
}

// delete conflict with delete in go
func deletion(pr **node, data Interface) Interface {
	var ret Interface

	r := *pr
	cmp := data.Compare(r.data)
	if cmp < 0 {
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
			return r.data
		}
		// something error
		if !isRed(r.right) && !isRed(r.right.left) {
			moveRedRight(pr)
			r = *pr
			cmp = data.Compare(r.data)
		}
		if cmp == 0 {
			ret = r.data
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
