package foundation

type LinkedList struct {
	next *LinkedList
	v    int
}

func NewNode(v int) *LinkedList {
	node := &LinkedList{v: v}
	return node
}

func ListSize(LinkedList LinkedList) int {
	count := 1
	for LinkedList.next != nil {
		count += 1
	}
	return count
}
