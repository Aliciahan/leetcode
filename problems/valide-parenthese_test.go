package problems

import "testing"

func TestParenthesesIsValid(t *testing.T) {

	a:= "()"
	b:= "()[]{}"
	c:= "(]"
	d:= "([)]"
	e:= "{[]}"

	if ParenthesesIsValid(a) != true {
		t.Error("A Failed")
	}
	if ParenthesesIsValid(b) != true {
		t.Error("B Failed")
	}
	if ParenthesesIsValid(c) != false {
		t.Error("C Failed")
	}
	if ParenthesesIsValid(d) != false {
		t.Error("D Failed")
	}
	if ParenthesesIsValid(e) != true {
		t.Error("E Failed")
	}
}
