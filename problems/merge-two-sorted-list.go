package problems

import (
	. "github.com/Aliciahan/leetcode/data-structure"
)

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
#we can not operate with l1 because that is to say l1 has been changed after operation.

we operate with the l1....as it mentioned in description
note that l2 is also sorted, so that, we need an anchor, while insertion.

*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	anchor := l1
	for ;l2!=nil; l2=l2.Next{
		newNode := &ListNode{
			Val:l2.Val,
			Next:nil,
		}

		if anchor.Val > l2.Val {
			//append left
			newNode.Next = anchor
			anchor = newNode
			l1 = newNode
		} else {
			for !( anchor.Next == nil || (anchor.Val <= l2.Val && anchor.Next.Val > l2.Val)) {
				anchor = anchor.Next
			}
			//append right
			newNode.Next = anchor.Next
			anchor.Next = newNode
			anchor = newNode
		}

	}
	return l1

}
