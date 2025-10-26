package easy

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "basic merge - interleaved elements",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "nums2 is empty",
			nums1:    []int{1, 2, 3},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "nums1 is empty (m=0)",
			nums1:    []int{0, 0, 0},
			m:        0,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "all nums1 elements are smaller",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "all nums2 elements are smaller",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "single element in each array",
			nums1:    []int{2, 0},
			m:        1,
			nums2:    []int{1},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "single element arrays - nums1 smaller",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{2},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "duplicate elements",
			nums1:    []int{1, 1, 1, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 1, 1},
			n:        3,
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:     "negative numbers",
			nums1:    []int{-3, -1, 0, 0, 0},
			m:        2,
			nums2:    []int{-2, 0, 1},
			n:        3,
			expected: []int{-3, -2, -1, 0, 1},
		},
		{
			name:     "only one element from nums2",
			nums1:    []int{1, 3, 5, 0},
			m:        3,
			nums2:    []int{2},
			n:        1,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "only one element from nums1",
			nums1:    []int{3, 0, 0, 0},
			m:        1,
			nums2:    []int{1, 2, 4},
			n:        3,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since merge modifies nums1 in place
			nums1Copy := make([]int, len(tt.nums1))
			copy(nums1Copy, tt.nums1)

			merge(nums1Copy, tt.m, tt.nums2, tt.n)

			if !reflect.DeepEqual(nums1Copy, tt.expected) {
				t.Errorf("merge() = %v, want %v", nums1Copy, tt.expected)
			}
		})
	}
}
