package easy

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	tests := []struct {
		name   string
		levels int
		want   [][]int
	}{
		{
			name:   "example 1",
			levels: 2,
			want:   [][]int{{1}, {1, 1}},
		},
		{
			name:   "example 2",
			levels: 3,
			want:   [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name:   "example 3",
			levels: 4,
			want:   [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.levels)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PascalTriangle(%v) = %v, want %v", tt.levels, got, tt.want)
			}
		})
	}
}

