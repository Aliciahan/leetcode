package problems

import "github.com/Aliciahan/leetcode/data-structure"

//https://blog.csdn.net/Try_harder_every_day/article/details/78687734


var (
	specialChars = []string{"+","-","*","/","(",")"}
)

func levelSpealChars(s string) int {
	switch s {
	case "+","-":
		return 0
	case "*","/":
		return 1
	case "(", ")":
		return 2
	}
	return 0
}


func containsSpecialChars(s string) bool{
	for _,i := range specialChars {
		if s == i {
			return true
		}
	}
	return false
}


func CutExpressionToStrings(exp string) [] string {
	var ret []string
	point := 0
	expBytes := []byte(exp)
	for i:=0;i<len(expBytes); i++ {
		if containsSpecialChars(string(expBytes[i])) {
			if point != i {
				ret = append(ret,string(expBytes[point:i]))
			}
			ret = append(ret,string(expBytes[i]))
			point = i+1
		}
	}
	if point != len(expBytes) {
		ret = append(ret,string(expBytes[point:len(expBytes)]))
	}
	return ret
}


func TransLateToRPN(original []string) [] string {

	var ret [] string
	s := data_structure.NewStackList()
	for _,i := range original {
		if !containsSpecialChars(i) {
			ret = append(ret,i)
		}else {
			if s.Empty() || levelSpealChars(i)>=levelSpealChars(s.Peak().(string)){
				s.Push(i)
			}else{
				ret = append(ret, s.Pop().(string))
			}
		}
	}
	for !s.Empty() {
		ret = append(ret, s.Pop().(string))
	}
	return ret

}
