package easy

import "math"

// 3354. Make Array Elements Equal to Zero
func countValidSelections(nums []int) int {
	fullSum := 0
	for i := range nums {
		fullSum += nums[i]
	}
	sumUntil := 0
	results := 0
	for i := range nums {
		sumUntil += nums[i]
		if nums[i] != 0 {
			continue
		}

		if fullSum-sumUntil*2 == 0 {
			results += 2
		} else if math.Abs(float64(fullSum-sumUntil*2)) == 1 {
			results += 1
		}
	}
	return results
}
