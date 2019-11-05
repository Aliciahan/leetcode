package problems

import (
	. "github.com/Aliciahan/leetcode/data-structure"
	"testing"
)

var linkedList1 = &ListNode{
	Val:  1,
	Next: &ListNode{
		Val:  2,
		Next: nil,
	},
}

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome(linkedList1) != false {
		t.Error("1-2 is not palindrome" )
	}
}
