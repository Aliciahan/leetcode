package problems

func CutBar(longer int, worker int, currentCut int) int {
	if currentCut >= longer {
		return 0 // if current cut pieces is equal or bigger than length of the bar, return 0
	} else if currentCut < worker {
		return CutBar(longer,worker,currentCut*2)+1
		// if current cut pieces is less than worker,
		// we cut one time, thus each worker has a bar to cut.
	} else {
		return CutBar(longer,worker,currentCut+worker)+1
		// if current cut pieces is more than workers,
		// each worker can cut one  bar in two, then the bar num will be current cut +worker number
	}
}


