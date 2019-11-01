package problems

func reverse(x int) int {

	flag := true
	if x == 0 || x>147483647 || x<(-147483648){
		return 0
	}else if x < 0 {
		flag = false
		x = -x
	}

	resultat := 0
	for ;x/10 >= 1; {
		resultat = resultat*10 + x%10
		x = x/10
	}

	resultat = resultat*10 + x%10

	if flag {
		return resultat
	}else {
		return 0-resultat
	}


}
