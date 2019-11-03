package problems

import data_structure "github.com/Aliciahan/leetcode/data-structure"

/*
1. string to list
2. stack FILO
3. if [ ( { push,
4. if ] ) }, verify pop() is the char correspond, if not false
5. at last, the stack should be vid, otherwise false
 */


func ParenthesesIsValid(s string) bool  {

	sAsSlice := []byte(s)
	stackFILO := data_structure.NewStackList()

	for _,i := range sAsSlice {
		charString := string(i)
		switch charString {
		case "[", "(", "{":
			stackFILO.Push(charString)
		case "]":
			if stackFILO.Pop() != "["{
				return false
			}
		case ")":
			if stackFILO.Pop() != "("{
				return false
			}
		case "}":
			if stackFILO.Pop() != "{"{
				return false
			}
		}
	}
	if stackFILO.Len() != 0 {
		return false
	}

	return true


}