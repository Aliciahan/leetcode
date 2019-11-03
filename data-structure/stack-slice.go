package data_structure

import (
	"fmt"
	"sync"
)

//https://hansedong.github.io/2019/04/02/15/

type Item interface{}

type ItemStackSlice struct {
	items [] Item
	lock sync.RWMutex
}

func NewStack() *ItemStackSlice {
	s:= &ItemStackSlice{}
	s.items = [] Item{}
	return s
}

func (s *ItemStackSlice) Print () {
	fmt.Println(s.items)
}

func (s *ItemStackSlice) Push (t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items,t)
}

func (s *ItemStackSlice) Pop() Item{
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items)==0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0:len(s.items)-1]
	return item
}

func (s *ItemStackSlice) Len() int{
	return len(s.items)
}

