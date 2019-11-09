package problems
type Job struct {
	id          int
	jobTime     int
	childJobIDs []int
}

/*
Recursive :
	reture = jobTime + children_jobTimes
	- if the childJobIDs len = 0 return
 */

/*
jobs list to LinkedList
 */
func FindTotalJobTime(j Job) int {
	jobs := []Job{
		Job{1, 30, []int{2, 4}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}

	if len(j.childJobIDs) == 0 {
		return j.jobTime
	} else {
		var sum int
		for _,i := range j.childJobIDs {
			for _,walk := range jobs {
				if walk.id == i {
					sum += FindTotalJobTime(walk)
				}
			}
		}
		return j.jobTime + sum
	}
}

//find_total_job_time(jobs, 1) # 30 + (10+20) + 60 = 120
//find_total_job_time(jobs, 2) # 10 + 20 = 30
//find_total_job_time(jobs, 4) # 60
