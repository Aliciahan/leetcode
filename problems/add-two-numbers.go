package problems

import (
	. "github.com/Aliciahan/leetcode/data-structure"
)

//You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example:
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.




func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result ListNode
    p := & result
    nextDeci := 0

    for ;l1.Next!=nil && l2.Next!=nil;  {

    	p.Val = l1.Val+l2.Val + nextDeci
    	if p.Val >=10 {
    		nextDeci = 1
    		p.Val = p.Val - 10
		} else {
			nextDeci = 0
		}

    	p.Next = &ListNode{
    		Val: 0,
    		Next: nil,
		}

    	p = p.Next
		l1=l1.Next
		l2=l2.Next
	}

	p.Val = l1.Val+l2.Val+nextDeci
	if l1.Next == nil {
		p.Next = l2.Next
	} else if l2.Next == nil {
		p.Next = l1.Next
	}

	for p.Val >=10 {
		p.Val = p.Val -10
		if p.Next == nil{
			p.Next = &ListNode{
				Val:0,
				Next: nil,
			}
		}
		p=p.Next
		p.Val = p.Val +1
	}


	return &result
}

