package easy

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want int
	}{
		{
			name:   "example 1",
			nums:   []int{1, 2, 3, 4, 5},
			target: 4,
			want: 3,
		},
		{
			name:   "example 2",
			nums:   []int{1,3,5,6},
			target: 7,
			want: 4,
		},
		{
			name:   "example 3",
			nums:   []int{1,3,5,6},
			target: 0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInsert(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
