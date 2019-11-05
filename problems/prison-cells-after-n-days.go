package problems

import "fmt"

func prisonAfterNDays(cells []int, N int) []int {
	ret := cells
	cycleSize := prisonDetectCircle(cells)
	for i:=0; i<N%cycleSize+cycleSize; i++{
		ret = prisonAfterOneDay(ret)
		fmt.Println("the current cells is: ", ret)
	}
	return ret
}

func prisonDetectCircle(cells []int) int{
	// 0 1 2 3 4 ; when 4 is the same with 1, the length of cycle is 4-1 =3
	statusPool := make(map[string]int)
	ret := cells
	statusPool[fmt.Sprint(ret)] = 0
	fmt.Println(statusPool)

	counter := 1
	ok := false
	v := 0
	for !ok {
		ret = prisonAfterOneDay(ret)
		key := fmt.Sprint(ret)
		v,ok = statusPool[key]
		fmt.Println("v is:", v)
		statusPool[key] = counter
		counter ++
		fmt.Println("update status pool:",statusPool)
	}
	return counter - v -1
}

func prisonAfterOneDay(cells []int) []int{
	length:=len(cells)
	var ret []int
	ret = append(ret, 0)
	for i:=1; i<length-1; i++ {
		if cells[i-1] == cells[i+1]{
			ret = append(ret, 1)
		}else{
			ret = append(ret, 0)
		}
	}
	ret = append(ret, 0)
	return ret
}

