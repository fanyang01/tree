package splay

// splay bottom-up, usually call it in zigRoot()
// cmpFlag will be passed to function 'compare'
func splay(r *node, data Interface, cmpFlag int) (*node, int) {
	var (
		cmp, path, prev int
	)
	if cmp = compare(data, r.data, cmpFlag); cmp == 0 {
		return r, -1
	} else if cmp < 0 {
		// sometimes data to access is not in tree
		if r.left == nil {
			return r, -1
		}
		r.left, prev = splay(r.left, data, cmpFlag)
		path = LEFT
	} else {
		if r.right == nil {
			return r, -1
		}
		r.right, prev = splay(r.right, data, cmpFlag)
		path = RIGHT
	}
	if prev == -1 {
		return r, path
	}
	path = prev<<1 | path
	return splayPath(r, path)
}

func splayPath(r *node, path int) (*node, int) {
	// determine path
	switch path {
	case LEFTLEFT:
		r = rightZigZig(r)
	case LEFTRIGHT:
		r = doubleLeftRotate(r)
	case RIGHTLEFT:
		r = doubleRightRotate(r)
	case RIGHTRIGHT:
		r = leftZigZig(r)
	}
	return r, -1
}

func zigRoot(r *node, path int) *node {
	switch path {
	case LEFT:
		return rightRotate(r)
	case RIGHT:
		return leftRotate(r)
	}
	return r
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

/*
 *            g
 *           / \
 *          p   d
 *         / \
 *        x   c
 *       / \
 *      a   b
 *  ->
 *          x
 *         / \
 *        a   p
 *           / \
 *          b   g
 *             / \
 *            c   d
 */
func rightZigZig(g *node) *node {
	p := g.left
	x := p.left
	g.left = p.right
	p.right = g
	p.left = x.right
	x.right = p
	return x
}

/*
 *          g
 *         / \
 *        a   p
 *           / \
 *          b   x
 *             / \
 *            c   d
 *  ->
 *            x
 *           / \
 *          p   d
 *         / \
 *        g   c
 *       / \
 *      a   b
 */
func leftZigZig(g *node) *node {
	p := g.right
	x := p.right
	g.right = p.left
	p.left = g
	p.right = x.left
	x.left = p
	return x
}
