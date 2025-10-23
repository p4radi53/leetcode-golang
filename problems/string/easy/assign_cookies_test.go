package string

import (
	"reflect"
	"testing"
)

func TestAssignCookies(t *testing.T) {
	tests := []struct {
		name    string
		greed   []int
		cookies []int
		want    int
	}{
		{
			name:    "example 1",
			greed:   []int{1, 2, 3},
			cookies: []int{1, 2, 3},
			want:    3,
		},
		{
			name:    "more children",
			greed:   []int{1, 2, 3, 4},
			cookies: []int{1, 2, 3},
			want:    3,
		},
		{
			name:    "more cookies",
			greed:   []int{1, 2, 3},
			cookies: []int{1, 2, 3, 4},
			want:    3,
		},
		{
			name:    "too much greed",
			greed:   []int{1, 2, 4, 5},
			cookies: []int{1, 2, 3, 4},
			want:    3,
		},
		{
			name:    "unsorted",
			greed:   []int{10, 9, 8, 7},
			cookies: []int{5, 6, 7, 8},
			want:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findContentChildren(tt.greed, tt.cookies)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContentChildren(%v, %v) = %v, want %v", tt.greed, tt.cookies, got, tt.want)
			}
		})
	}
}
