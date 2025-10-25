package easy

func twoSum(nums []int, target int) []int {
	alreadyEncountered := make(map[int]int)

	for ix, val := range nums {
		pair := target - val
		if x, exists := alreadyEncountered[pair]; exists {
			return []int{x, ix}
		} else {
			alreadyEncountered[val] = ix
		}
	}

	return []int{0, 0}
}
