package easy

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "example 1",
			input: 1,
			want:  1,
		},
		{
			name:  "example 1",
			input: 4,
			want:  5,
		},
		{
			name:  "example 1",
			input: 5,
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.input)
			if got != tt.want {
				t.Errorf("ClimbStairs(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
