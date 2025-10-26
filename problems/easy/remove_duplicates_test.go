package easy

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "example 1",
			nums: []int{1, 1, 2, 3},
			want: 4,
		},
		{
			name: "example 1",
			nums: []int{1, 1, 2, 2, 3},
			want: 5,
		},
		{
			name: "example 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
		{
			name: "example 1",
			nums: []int{1, 1, 1, 2, 2, 3, 3, 3},
			want: 6,
		}, {
			name: "example 1",
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
