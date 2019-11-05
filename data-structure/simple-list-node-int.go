package data_structure

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) PrintAll(){
	po := l
	for ; po.Next!= nil ;po=po.Next {
		fmt.Print(po.Val)
	}
	fmt.Print(po.Val)
	fmt.Print("\n")
}



