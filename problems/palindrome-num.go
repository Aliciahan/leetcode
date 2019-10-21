package problems

import (
	"fmt"
	"strings"
)

func ReverseString(s string) string {
	ss := strings.Split(s,"")
	l := len(ss)
	if len(ss) >1 {
		return ss[l-1] + ReverseString(strings.Join(ss[:l-1],""))
	} else {
		return ss[0]
	}
}


func palindromeWithBase(i int, typeInput string)bool{
	var s string
	switch typeInput {
	case "2":
		s = fmt.Sprintf("%b",i)
	case "8":
		s = fmt.Sprintf("%o",i)
	case "10":
		s = fmt.Sprintf("%d",i)
	default:
		return false
	}

	if s == ReverseString(s) {
		return true
	}
	return false
}

func GetSmallestpalindromebase28(startFrom int){

	stopFlag := false
	for ;!stopFlag; startFrom++ {
		if startFrom%2 != 1 {
			continue
			//for the binary number, it can not be 0 as top, thus, 1, we can omit all the doubles
		}
		if palindromeWithBase(startFrom,"2") && palindromeWithBase(startFrom, "8") && palindromeWithBase(startFrom, "10"){
			stopFlag = true
			fmt.Println("the result is : ", startFrom)
		}
	}
}
