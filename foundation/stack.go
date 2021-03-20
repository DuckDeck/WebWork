package foundation

import (
	"fmt"
)

type Element interface{}

var header *entry
var minEntry *entry
var size int

type entry struct {
	previous *entry
	next     *entry
	element  Element
}

func newEntry(prev, next *entry, e Element) *entry {
	return &entry{prev, next, e}
}

func NewStack() *entry {
	header = newEntry(nil, nil, nil)
	header.previous = header
	header.next = header
	return header
}

type Stack interface {
	Push(e Element)
	Pop() Element
	Top() Element
	Clear() bool
	Size() int
	IsEmpty() bool
	MinValue() Element
}

func (e *entry) Push(element Element) {
	addBefore(header, element)
}

func (e *entry) Pop() Element {
	if e.IsEmpty() {
		fmt.Printf("Stack is empty")
		return nil
	}
	return header.previous.element
}

func (e *entry) Clear() bool {
	if e.IsEmpty() {
		fmt.Printf("Stack is empty!")
		return false
	}
	entry := header.next
	for entry != header {
		nextEntry := entry.next
		entry.next = nil
		entry.previous = nil
		entry.element = nextEntry.IsEmpty
		entry = nextEntry
	}
	header.next = header
	header.previous = header
	size = 0
	return true
}

func (e *entry) Size() int {
	return size
}

func (e *entry) IsEmpty() bool {
	if size == 0 {
		return true
	}
	return false
}

func addBefore(e *entry, element Element) Element {
	newEntry := newEntry(e.previous, e, element)
	newEntry.previous.next = newEntry
	newEntry.next.previous = newEntry
	size++
	return newEntry
}

//用两个栈实现队列
//我们有两个栈，S1，S2。 S1只用来进队列，S2只用来出队列
//入队时，判断S1能不能放，不能放的话，判断S2是不是空，如果是空的就把S1的数据全部倒入S2，再将新的元素放入S1，
//如果S2不是空的说明满了需要把S2清空才能手队列，如果S1能放就直接放
//出队列很简单，直接判断S2是不是空，如果不是空就弱出，是空的话就奖S1的全部数据倒入S2，再将顶部数据弱出

//代码...
//这个比较麻烦

//输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
