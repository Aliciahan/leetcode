package problems

import "testing"

func TestCutBar(t *testing.T) {
	re :=CutBar(20,3,1)
	if re == 8 {
		t.Log("Success")
	}else{
		t.Error("Failed, the result should be 8, but :", re )
	}
}
