package problems

import (
	"testing"
)

func TestGetSmallestpalindromebase28(t *testing.T){
	GetSmallestpalindromebase28(10)
}

func TestReverseString(t *testing.T) {
	if ReverseString("12345") != "54321" {
		t.Error("Failed")
	}
}
