package foundation

type Tree struct {
	v     int
	left  *Tree
	right *Tree
}

func CreatesSortTree(tree Tree, v int) {
	if tree.left == nil && tree.right == nil {
		tree.v = v
	}

}

//判断是不是镜像树
func IsMirrorTree(tree *Tree) bool {
	if tree == nil {
		return true
	}
	return isMirror(tree.left, tree.right)
}
func isMirror(left *Tree, right *Tree) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.v != right.v {
		return false
	}
	ret1 := isMirror(left.left, right.right)
	ret2 := isMirror(left.right, right.left)
	return ret1 && ret2
}

func TestMirrorTree() bool {
	tree := Tree{v: 0}
	tree.left = &Tree{v: 12}
	tree.right = &Tree{v: 12}
	tree.left.left = &Tree{v: 15}
	tree.right.right = &Tree{v: 15}
	return IsMirrorTree(&tree)
}
