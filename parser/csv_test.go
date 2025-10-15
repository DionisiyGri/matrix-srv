package parser

import (
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	parser := New()

	tests := []struct {
		name     string
		input    string
		expected [][]string
		hasError bool
	}{
		{
			name:     "Valid data in csv matrix",
			input:    "1,2,3\n4,5,6\n7,8,9\n",
			expected: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			hasError: false,
		},
		{
			name:     "Empty matrix",
			input:    "",
			expected: [][]string{},
			hasError: false,
		},
		{
			name:     "Invalid matrix format",
			input:    "1,2,3\n4,5\n7,8,9\n", // Missing value in the second row
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := parser.ParseCSV(strings.NewReader(tt.input))
			if (err != nil) != tt.hasError {
				t.Fatalf("Expected error: %v, but got: %v", tt.hasError, err)
			}
			if !equal(output, tt.expected) {
				t.Errorf("Expected output: %v, but got: %v", tt.expected, output)
			}
		})
	}
}
