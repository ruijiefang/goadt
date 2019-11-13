package storage

//
// Implementation of LLRB.
// Ref: https://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//

// definitions of red and black.
const (
	Red = iota
	Black = iota
)

// a single node inside the tree.
type tree_node struct {
	left *tree_node
	right *tree_node
	color int
	key int
	val interface{}

}

// RedBlackTree ADT.
type RedBlackTree struct {
	root *tree_node
	size uint32
}

// left rotation
func rotate_left(parent *tree_node) {
	var right_child = parent.right
	parent.right = right_child.left
	right_child.left = parent
	right_child.color = parent.color
	parent.color = Red
}

// right rotation
func rotate_right(parent *tree_node) {
	var left_child = parent.left
	parent.left = left_child.right
	left_child.right = parent
	left_child.color = parent.color
	parent.color = Red
}

// flip the color of a single node
func color_flip_1(h *tree_node) {
	if h.color == Red {
		h.color = Black
	} else {
		h.color = Red
	}
}

// flip the color of a parent and its children
func color_flip(h *tree_node) {
	color_flip_1(h)
	if (h.left != nil) {
		color_flip_1(h.left)
	}
	if (h.right != nil) {
		color_flip_1(h.right)
	}
}

func move_red_left(h *tree_node) *tree_node {
	color_flip(h)
	if //TODOTODO
}

func move_red_right(h *tree_node) *tree_node {

}
func (t *RedBlackTree) Size() uint32 {
	return t.size
}

// Returns a new RedBlackTree type
func New() *RedBlackTree {
	return &RedBlackTree{root: nil, size:0}
}

// Existence
func (t *RedBlackTree) Exists(key int) bool {
	var ptr = t.root
	for ptr != nil {
		if ptr.key == key {
			return true
		} else if ptr.key < key {
			ptr = ptr.left
		} else {
			ptr = ptr.right
		}
	}
	return false
}

func (t *RedBlackTree) Find(key int) interface{} {
	var ptr = t.root
	for ptr != nil {
		if ptr.key == key {
			return ptr.val
		} else if ptr.key < key {
			ptr = ptr.left
 		} else /* ptr.key > key */ {
 			ptr = ptr.right
		}
	}
	return nil
}

func insert(h * tree_node, key int, val interface{}) *tree_node {
	if h == nil {
		return &tree_node{
			left:  nil,
			right: nil,
			color: Red,
			key:   key,
			val:   nil,
		}
	}

	if h.left.color == Red && h.right.color == Red {
		color_flip(h)
	}

	if h.key == key {
		h.val = val
	} else if h.key < key {
		h.left = insert(h.left, key, val)
	} else /* h.key > key */ {
		h.right = insert(h.right, key, val)
	}
}

func (t *RedBlackTree) Insert(key int, val interface{}) {
	if t.size == 0 {
		t.root = &tree_node{
			left:  nil,
			right: nil,
			color: Red,
			key:   key,
			val:   val,
		}
		return
	}
	t.root = insert(t.root, key, val)
	t.root.color = Black
}

func delete_min(h *tree_node) *tree_node {
	if h.left == nil {
		return nil
	}

	if !(h.left.color == Red) && !(h.left.left.color == Red) {
		h =
	}
}

func delete(t *RedBlackTree, key int, val interface{}) {

}

func (t *RedBlackTree) Delete(key int, val interface{}) {
	if t.size == 0 {
		t.root = &tree_node{
			left:  nil,
			right: nil,
			color: Red,
			key:   key,
			val:   val,
		}
		return
	}
	delete(t, key, val)
	t.root.color = Black
}