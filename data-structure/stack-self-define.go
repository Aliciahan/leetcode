package data_structure

import "sync"

type(
	StackSelfDefined struct {
		top *node
		length int
		lock *sync.RWMutex
	}
	node struct {
		value interface {}
		prev *node
	}
)


func NewStackSelfDefined() * StackSelfDefined {
	return &StackSelfDefined{nil, 0,&sync.RWMutex{}}
}


func (t *StackSelfDefined) Len() int {
	return t.length
}

func (t *StackSelfDefined) Peak() interface{}{
	if t.length == 0 {
		return nil
	}
	return t.top.value
}


func (t *StackSelfDefined) Pop() interface{}{
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.length == 0 {
		return nil
	}

	n := t.top
	t.top=t.top.prev
	t.length --
	return n.value
}

func (t * StackSelfDefined) Push(value interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	n := &node{
		value,
		t.top,
	}
	t.top = n
	t.length++

}
