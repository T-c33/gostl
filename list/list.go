package list

import "github.com/pkg/errors"

type Node struct {
	value interface{}
	next *Node
}

func (node *Node) GetNodeValue() interface{} {
	return node.value
}

func (node *Node) GetNodeNext() *Node {
	return node.next
}

type List struct {
	head *Node
	tail *Node
	length int
}

func (l *List) GetListHead() *Node {
	return l.head
}

func (l *List) GetListTail() *Node {
	return l.tail
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

func (l *List) Insert(value interface{}, position int) (err error){
	if position > (*l).length || (*l).length == 0 {
		return errors.New("position err or list length error")
	}
	node := NewNode(value)
	if position == 0 {
		(*node).next = (*l).head
		(*l).head = node
	} else {
		preItem := (*l).head
		for i := 1; i < position; i++ {
			preItem = (*preItem).next
		}
		(*node).next = (*preItem).next
		(*preItem).next = node
	}
	(*l).length++
	if position == (*l).length {
		(*l).tail = node
	}
	return nil
}

func (l *List) GetNode(position int) (node *Node) {
	if position >= l.length {
		return nil
	}
	node = (*l).head
	for i := 0; i < position; i++ {
		node = (*node).next
	}
	return node
}

func (l *List) GetValue(position int) (value interface{}) {
	node := l.GetNode(position)
	if node == nil {
		return nil
	} else {
		return node.GetNodeValue()
	}
}

func (l *List) Remove(position int, node *Node) (err error) {
	if position >= (*l).length {
		return errors.New("wrong position")
	}

	if position == 0 {
		node = (*l).head
		(*l).head = (*node).next
		if (*l).length == 1 {
			(*l).tail = nil
		}
	} else {
		preItem := (*l).head
		for i := 1; i < position; i++ {
			preItem = (*preItem).next
		}

		node = (*preItem).next
		(*preItem).next = (*node).next

		if position == ((*l).length - 1) {
			(*l).tail = preItem
		}
	}
	(*l).length--
	return nil
}
