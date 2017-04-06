package tree

type color bool

const (
	red   color = false
	black       = true
)

func (n *node) isBlack() bool {
	return (n.payload.(color) == black)
}

func (n *node) ancestor(i int) *node {
	for j := 0; j < i; j++ {
		if n == nil {
			return n
		}
		n = n.parent
	}
	return n
}

func rbSearch(n *node) {
	//Todo
}
func rbInsert(n *node) {
	// If i is the root
	for {
		p := n.parent
		if p == nil {
			n.payload = black
			return
		}
		// i's parent must exist, as i is not the root ---
		// If i's parent is black
		if p.isBlack() {
			return
		}

		// i's grandparent must exist, as i's parent is red. ---
		// if i's grandparent did not exist, i's parent would
		// be the root and would be black.
		// If i's parent is red and i's uncle is red

		redUncle := false
		gp := p.parent
		if gp.left == p {
			if !gp.right.isBlack() {
				redUncle = true
			}
		} else if !gp.left.isBlack() {
			redUncle = true
		}
		// Would this be faster?
		// if bst.color(right(gp)) == RED && bst.color(left(gp)) == RED {...}
		if redUncle {
			n = gp
			gp.left.payload = black
			gp.right.payload = black
			gp.payload = red
			// recurse on i's grandparent

		} else {
			gp.payload = red
			if gp.left == p {
				// if i is a right child
				if n.parent.left == n {
					n.parent.leftRotate()
				}
				n.parent.payload = black
				n.ancestor(2).rightRotate()
			} else {
				if n.parent.left == n {
					n.parent.rightRotate()
				}
				n.parent.payload = black
				n.ancestor(2).leftRotate()
			}
			return
		}
	}
}
func rbDelete(n *node) {
	// If this node has two children
	if n.right != nil && n.left != nil {
		// Find the maximum value of the left subtree
		// or the minimum value of the right subtree.
		// Presumably defaulting to one over the other will
		// cause the tree to lean in one direction over the
		// other. Needs to be tested.

		// if rand.Float64() < 0.5 {
		n2 := n.right.minKey()
		//} else {
		// n2 := n.left.maxKey()
		//}

		n2.deleteSwap(n)
	}
	//Todo: the difficult cases
}