package easy

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
