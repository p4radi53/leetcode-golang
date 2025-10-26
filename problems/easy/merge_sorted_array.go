package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	currentIndex := m + n - 1

	i, j := m-1, n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[currentIndex] = nums1[i]
			i--
		} else {
			nums1[currentIndex] = nums2[j]
			j--
		}
		currentIndex--
	}

	// after first array empties
	// fill the beginning with the rest of the second array
	for j >= 0 && i < 0 {
		nums1[currentIndex] = nums2[j]
		currentIndex--
		j--
	}
}
