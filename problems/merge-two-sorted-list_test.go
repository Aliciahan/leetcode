package problems

import (
	. "github.com/Aliciahan/leetcode/data-structure"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:4,
				Next: nil,
			},
		},
	}
	l1.PrintAll()
	l2.PrintAll()

	l3 := mergeTwoLists(l1,l2)
	l3.PrintAll()

}
