package problems

import (
	"fmt"
	"testing"
)

func TestPrisonAfterNDays(t *testing.T) {
	t1 := []int {0,1,0,1,1,0,0,1}
	ret := prisonAfterNDays(t1,7)
	fmt.Println("Result is :", ret)

	fmt.Println("length of cycle:",prisonDetectCircle(t1))

}
