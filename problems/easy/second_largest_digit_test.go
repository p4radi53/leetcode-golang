package easy

import "testing"

func TestSecondHighest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "basic case with letters and digits",
			input:    "dfa12321afd",
			expected: 2,
		},
		{
			name:     "only one unique digit",
			input:    "abc1111",
			expected: -1,
		},
		{
			name:     "no digits",
			input:    "abcdef",
			expected: -1,
		},
		{
			name:     "all digits 0-9",
			input:    "0123456789",
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := secondHighest(tt.input)
			if result != tt.expected {
				t.Errorf("secondHighest(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
