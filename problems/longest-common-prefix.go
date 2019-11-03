package problems

func LongestCommonPrefix(strs []string) string {

	longestCounter := 0

	lengthStrs := 0
	var byteStrs [][]byte
	for _, i := range strs {
		lengthStrs+=1
		byteStrs = append(byteStrs, []byte(i))
	}
	if lengthStrs == 0 {
		return ""
	}

	flag := true
	for i:=0; i<len(byteStrs[0]) && flag; {
		for j := 1; j<lengthStrs; j++{
			if len(byteStrs[j])-1 < i || byteStrs[0][i] != byteStrs[j][i]{
				flag = false
				return string(byteStrs[0][0:longestCounter])
			}
		}
		longestCounter ++
		i++
	}
	return string(byteStrs[0][0:longestCounter])
}
