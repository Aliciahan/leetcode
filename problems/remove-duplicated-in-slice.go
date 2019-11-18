package problems

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/

func removeDuplicates(nums []int) int {

	for i:=0;i<len(nums)-1;{
		if nums[i] == nums[i+1] {
			nums = append(nums[:i],nums[i+1:]...)
		} else{
			i++
		}

	}
	return len(nums)

}
