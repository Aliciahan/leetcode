package problems

import (
	"math"
)

/*
Link to the pb description: https://leetcode.com/problems/flip-string-to-monotone-increasing/

case study: if len(s) = 5, we have the possibilities of 00000 00001 00011 00111 01111 11111, totally 6.
Then, for each case, we have 0 s in the left, 1 in the right.
	we should calculate how many 1 in left, how many 0 in right.

PREFIX SUM:

		|s[0]| s[1]| s[2]| s[3]| s[4]|
p[i]	0    1     2     3     4     5
for example: S = "010110" will give us
      0 1 0 1 1 0
P = [0,0,1,1,2,3,3] that is : in the position p[i] how many 1 we have in front.
how to calculate how many 0 the right side? at position x, N-x-(p[N]-p[x]) will be the answer.

Thus for one solution, the number of flips is num(turn the left to 0) + num(turn the right to 1)

at last MIN(solutions)
*/

func MinFlipsMonoIncr(S string) int {
	// first we calculate the Prefix Sum List
	s := []byte(S)
	n := len(s)
	prefixSumList:= []int{0}
	for i:=0 ; i< n ;i++ {
		if  string(s[i]) == "1" {
			prefixSumList = append(prefixSumList,prefixSumList[i]+1)
		}else{
			prefixSumList = append(prefixSumList,prefixSumList[i])
		}
	}
	ret := math.MaxInt32
	for j:=0; j<=n; j++ {
		if ret >= prefixSumList[j]+n-j-(prefixSumList[n]-prefixSumList[j]) {
			ret = prefixSumList[j]+n-j-(prefixSumList[n]-prefixSumList[j])
		}
	}
	return ret
}
