package data_structure

import "testing"

var stack *ItemStackSlice

func init(){
	stack = NewStack()
}

func BenchmarkItemStackSlice_Push(b *testing.B) {
	for i:=0;i<b.N; i++ {
		stack.Push("test")
	}
}

func BenchmarkItemStackSlice_Pop(b *testing.B) {
	b.StopTimer()
	for i:=0; i<b.N; i++ {
		stack.Push("test")
	}
	b.StartTimer()
	for i:=0; i<b.N; i++ {
		stack.Pop()
	}
}
