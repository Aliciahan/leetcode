package problems

import "testing"
var(
	l1 = ListNode{
		Val: 2,
		Next: & ListNode{
			Val: 4,
			Next: & ListNode{
				Val:3,
				Next: nil,
			},
		},
	}

	l2 = ListNode{
		Val: 5,
		Next: & ListNode{
			Val: 6,
			Next: & ListNode{
				Val:4,
				Next: nil,
			},
		},
	}

	l3 = ListNode{
		Val: 5,
		Next: & ListNode{
			Val: 6,
			Next: & ListNode{
				Val: 7,
				Next: &ListNode{
					Val:9,
					Next: nil,
				},
			},
		},
	}


)
func TestAddTwoNumbers(t *testing.T) {

	res := addTwoNumbers(&l1,&l2)
	l1.PrintAll()
	l2.PrintAll()
	res.PrintAll()
	println("=================")
	////////
	l2.PrintAll()
	l3.PrintAll()
	println("L2+L3:")
	res2 := addTwoNumbers(&l2,&l3)
	res2.PrintAll()

}