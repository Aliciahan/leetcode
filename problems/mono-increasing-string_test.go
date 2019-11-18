package problems

import "testing"

func TestMinFlipsMonoIncr(t *testing.T) {
	t1 :="00110"
	t2 :="010110"
	t3 :="00011000"

	a1 := 1
	a2 := 2
	a3 := 2

	a1Ret := MinFlipsMonoIncr(t1)
	a2Ret := MinFlipsMonoIncr(t2)
	a3Ret := MinFlipsMonoIncr(t3)

	if a1Ret != a1 {
		t.Error("Err T1: the result is:",a1Ret, "should be:", a1 )
	}
	if a2Ret != a2 {
		t.Error("Err T2: the result is:",a2Ret, "should be:", a2 )
	}
	if a3Ret != a3 {
		t.Error("Err T3: the result is:",a3Ret, "should be:", a3 )
	}
}
