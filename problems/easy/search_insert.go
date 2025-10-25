package easy


// TODO
func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	low := 0
	high := len(nums) - 1

	middle := low + (high-low)/2
	middleValue := nums[middle]

	for low < high {
		if middleValue > target {
			high = middle - 1
		} else if middleValue < target {
			low = middle + 1
		} else {
			return middle
		}
		middle = low + (high-low)/2
		middleValue = nums[middle]
	}

	return middle + 1
}
