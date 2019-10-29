package problems


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