package problems

import "testing"

func TestReturnChange(t *testing.T) {

	coinValues := []int {1,3,5}

	t1 := ReturnChange(coinValues,3)
	if( t1 != 1 ){
		t.Error("Retuen should be 1 but :",)
	}

}
