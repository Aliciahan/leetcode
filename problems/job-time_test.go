package problems

import (
	"fmt"
	"testing"
)

func TestFindTotalJobTime(t *testing.T) {

	jobs := []Job{
		Job{1, 30, []int{2, 4}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}
	fmt.Println(FindTotalJobTime(jobs[0]))
}
