package data_structure

import (
	"container/list"
	"sync"
)


type StackList struct {
	list *list.List
	lock *sync.RWMutex
}


func NewStackList() * StackList {
	list := list.New()
	l := &sync.RWMutex{}
	return &StackList{
		list,
		l,
	}
}


func (stack *StackList) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}


func (stack *StackList) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	lastElement := stack.list.Back()
	if lastElement != nil {
		stack.list.Remove(lastElement)
		return lastElement.Value
	}
	return nil
}


func (stack *StackList) Peak() interface{}{
//view the top element
	lastElement := stack.list.Back()
	if lastElement != nil {
		return lastElement.Value
	}
	return nil
}


func (stack *StackList) Len() int {
	return stack.list.Len()
}


func (stack * StackList) Empty() bool {
	return stack.list.Len() == 0
}


