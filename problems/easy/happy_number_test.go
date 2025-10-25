package easy

import (
	"testing"
)

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "example 1",
			input: 19,
			want:  true,
		},
		{
			name:  "example 1",
			input: 2,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.input)
			if got != tt.want {
				t.Errorf("IsHappy(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
