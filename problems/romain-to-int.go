package problems

import "fmt"

func RomanToInt(s string) int {

	ret := 0

	valueMap := make(map[string]int)
	valueMap["I"] = 1
	valueMap["V"] = 5
	valueMap["X"] = 10
	valueMap["L"] = 50
	valueMap["C"] = 100
	valueMap["D"] = 500
	valueMap["M"] = 1000
	valueMap["IV"] = 4
	valueMap["IX"] = 9
	valueMap["XL"] = 40
	valueMap["XC"] = 90
	valueMap["CD"] = 400
	valueMap["CM"] = 900

	sByteSlice := []byte(s)
	length := len(sByteSlice)
	i:=0
	for ; i<length-1;  {
		windowSize2 := string(sByteSlice[i:i+2])
		windowSize1 := string(sByteSlice[i:i+1])
		if val,ok :=valueMap[windowSize2];ok {
			ret+=val
			i+=2
		}else{
			ret += valueMap[windowSize1]
			i+=1
		}
		fmt.Println("i := ", i)
	}
	if i >= length {
		return ret
	}else{
		return ret+valueMap[ string(sByteSlice[length-1:length]) ]
	}


}
