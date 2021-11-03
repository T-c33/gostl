package stack

import "sync"

type Item *interface{}
type Stack struct {
	items []Item
	lock sync.RWMutex
}

func (stack *Stack) Len() int {
	return len(stack.items)
}

func (stack *Stack) Cap() int {
	return cap(stack.items)
}

func NewStack() *Stack {
	s := &Stack{}
	s.items = []Item{}
	return s
}

func (stack *Stack) Push(t Item) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.items = append(stack.items, t)
}

func (stack *Stack) Pop() Item {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	length := len(stack.items)
	if length == 0 {
		return nil
	}
	item := stack.items[length - 1]
	stack.items = stack.items[:length - 1]
	return item
}
