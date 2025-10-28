package easy

import "testing"

func TestCountValidSelections(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "first example",
			nums: []int{1,0,2,0,3},
			want: 2,
		},
		{
			name: "second example",
			nums: []int{2,3,4,0,4,1,0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countValidSelections(tt.nums)
			if got != tt.want {
				t.Errorf("countValidSelections() = %v, want %v", got, tt.want)
			}
		})
	}
}

