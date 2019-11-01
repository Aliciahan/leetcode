package problems

import (
	"log"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func IsPalindrome(head *ListNode) bool {
	length := 1
	h := head
	tail := &ListNode{
		Next: nil,
		Val : head.Val,
	}
	h = h.Next
	for ;h!=nil; {
		length++
		t := tail
		tail = &ListNode{
			Val: h.Val,
			Next: t,
		}
		h = h.Next
	}
	log.Println("the length is:", length)
	for i:=0;i<length;i++ {
		if head.Val != tail.Val {
			log.Println("head:",head.Val," tail:", tail.Val)
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true

}
