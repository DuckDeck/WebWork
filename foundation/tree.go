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
