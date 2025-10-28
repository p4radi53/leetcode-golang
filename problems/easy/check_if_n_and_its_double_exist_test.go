package easy

import "testing"

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "example with double exists",
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			name: "example with no double",
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
		{
			name: "negative numbers with double",
			arr:  []int{7, 1, 14, 11},
			want: true,
		},
		{
			name: "array with zero",
			arr:  []int{0, 0},
			want: true,
		},
		{
			name: "single zero",
			arr:  []int{0},
			want: false,
		},
		{
			name: "negative and positive",
			arr:  []int{-2, 0, 10, -19, 4, 6, -8},
			want: false,
		},
		{
			name: "double at beginning",
			arr:  []int{2, 1, 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkIfExist(tt.arr)
			if got != tt.want {
				t.Errorf("checkIfExist(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
