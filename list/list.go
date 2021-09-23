package list

type Node struct {
	value interface{}
	next *Node
}

type List struct {
	head *Node
	tail *Node
	length int
}

func NewNode(value interface{}) *Node {
	n := &Node{}
	n.value = value
	n.next = nil
	return n
}

func NewList() *List {
	l := &List{}
	l.length = 0
	l.head = nil
	l.tail = nil
	return l
}

func (l *List) Len() int {
	return l.length
}

func (l *List) Append(value interface{}) {
	node := NewNode(value)
	(*node).next = nil
	if (*l).length == 0 {
		(*l).head = node
	} else {
		oldTail := (*l).tail
		(*oldTail).next = node
	}
	(*l).tail = node
	(*l).length++
}



