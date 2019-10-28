package problems

import (
	"fmt"
)

/*

Dynamic Programming

重点就是：
	1 大问题规模变小
	2 小问题和大问题同质， 只是规模区别

比如 硬币有1，3，5 凑够11元
i 为最终要凑够的钱数
j 为 需要的硬币数
我们理想的情况是 d(i) = j

当i=0， 1,3,5 都大于0，d(0) = 0
当i=1, 1<=i, d(1) = 1 + d(0) 我们拿一个硬币， 然后凑0元
when i = 2, 1<=i d(2) = 1+d(1) = 1+1+d(0) = 2
when i = 3 things become interesting: two choices:
	d(3) = 1 + d(0) // when we take 1 piece coin = 3
	d(3) = 1 + d(2) // when 3 times of 1
summary, d(3) should be d(3) = min(d(3-1)+1, d(3-3)+1)

In this way, the pb can be presented as :

i = the final sum
j = how many coin
v = the denomination of coin

d(i) = min{d(i-v*j) + 1}

 */


// values : coin values
// total: the money should be return
// s : the number of pieces (smallest)
func ReturnChange(values []int, total int) int {
	var s []int
	s = append(s,0)
	for cents:=1; cents <= total; cents ++ {
		minCoins := cents // when use the smallest value, we need a large num of coins
		for _,kind := range values {
			if kind <= cents {
				//if the cents is smaller than this kind of coin, check the s table
				//log.Println("cents:", cents," kinds:",kind)
				tmp := s[cents-kind]+1
				if tmp < minCoins {
					minCoins = tmp
				}
			}
		}
		s = append(s,minCoins)
		fmt.Println("value : ",cents," with num coins:", s[cents])
	}
	return s[total]
}

