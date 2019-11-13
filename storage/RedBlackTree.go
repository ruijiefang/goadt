package storage

const (
	Red = iota
	Black = iota
)

type tree_node struct {
	left *tree_node
	right *tree_node
	color int
	key int
	val interface{}

}

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

func rotate_right(parent *tree_node) {
	var left_child = parent.left
	parent.left = left_child.right
	left_child.right = parent
	left_child.color = parent.color
	parent.color = Red
}

func color_flip_1(h *tree_node) {
	if h.color == Red {
		h.color = Black
	} else {
		h.color = Red
	}
}

func color_flip(h *tree_node) {
	color_flip_1(h)
	color_flip_1(h.left)
	color_flip_1(h.right)
}

func (t *RedBlackTree) Size() uint32 {
	return t.size
}

func New() *RedBlackTree {
	return &RedBlackTree{root: nil, size:0}
}

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

func insert(h * tree_node, key int, val interface{}) {

}

func (t *RedBlackTree) Insert(key int, val interface{}) {
	if t.size == 0 {
		t.root = &tree_node{
			left:  nil,
			right: nil,
			color: Black,
			key:   key,
			val:   val,
		}
		return
	}
	insert(t.root, key, val)
	t.root.color = Black
}

func delete(t *RedBlackTree, key int, val interface{}) {

}

func (t *RedBlackTree) Delte(key int, val interface{}) {
	if t.size == 0 {
		t.root = &tree_node{
			left:  nil,
			right: nil,
			color: Black,
			key:   key,
			val:   val,
		}
		return
	}
	delete(t, key, val)
	t.root.color = Black
}