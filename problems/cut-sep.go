package problems

import (
	"fmt"
	"strings"
)

/*
  for each string in strings, we should distinguish items like host:a ; the sepreater is ","
  we use map[string]int
    if not exist, create entry
    else +1
  output map
*/
func CuttingStrings() {

	stream := []string{
		"host:a,role:web,availability-zone:us-east-1a",
		"host:b,role:web,availability-zone:us-east-1b",
		"host:a,role:web,availability-zone:us-east-1a",
		"host:c,role:db,db_role:master,availability-zone:us-east-1e",
		"host:d,role:db,db_role:replica,availability-zone:us-east-1a",
		"host:e,role:intake,availability-zone:us-east-1a",
		"host:e,role:intake,availability-zone:us-east-1a",
		"host:f,role:kafka,availability-zone:us-east-1a",
	}

	ret := make(map[string]int)

	for _,s := range stream {
		parts := strings.Split(s,",")
		for _,p := range parts {
			if _,ok := ret[p];ok {
				ret[p] += 1
			}else {
				ret[p] = 1
			}
		}
	}

	for key,item := range ret{
		fmt.Println(key,item)
	}
}
