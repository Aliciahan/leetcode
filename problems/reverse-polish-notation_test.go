package problems

import (
	"log"
	"testing"
)

func TestCutExpressionToStrings(t *testing.T) {
	l := CutExpressionToStrings("3+45*5")
	log.Println("l: ",l, " len: ", len(l))
	l2 := CutExpressionToStrings("5+(6-7)")
	log.Println("l2: ",l2, " len: ", len(l2))
}



func TestTransLateToRPN(t *testing.T) {
	s := "1+4+55*33"
	sArray := CutExpressionToStrings(s)
	result := TransLateToRPN(sArray)
	log.Println("s: ",s)
	log.Println("sRpn: ",result)
}