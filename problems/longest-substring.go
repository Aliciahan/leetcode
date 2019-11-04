package problems

func LengthOfLongestSubstringMap(original string) int{
	s := []byte(original)
	ans := 0
	n := len(s)
	mapping := make(map[byte]int)
	i := 0 //try to extend range [i,j]
	j:=0
	for  ; j< n;j++ {
		if charJ,contains := mapping[s[j]]; contains {
			if charJ > i {
				i = charJ
			}
		}
		if ((j-i)+1) > ans {
			ans = j-i+1
		}
		mapping[s[j]]=j+1
	}
	return ans
}

func LengthOfLongestSubstring(original string) int {
	biggestSize := 1
	windowSize := 1
	s := []byte(original)
	for i:=0; i<len(s)&&(i+windowSize)<len(s); {
		if checkRepeated(s[i:i+windowSize],s[i+windowSize]){
			i++
			windowSize = 1
		}else{
			windowSize++
			if windowSize > biggestSize{
				biggestSize = windowSize
			}
		}
	}

	return biggestSize
}

func checkRepeated(s []byte, new byte) bool {
	//return true when there are repeated
	for _, i := range s {
		if i == new {
			return true
		}
	}
	return false
}