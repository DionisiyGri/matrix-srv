package matrix

import (
	"reflect"
	"testing"
)

func TestEcho(t *testing.T) {
	m := NewCSVMatrixer()

	tests := []struct {
		name     string
		input    [][]string
		expected string
	}{
		{
			name: "Multiple rows",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
				{"10", "11", "12"},
			},
			expected: "1,2,3\n4,5,6\n7,8,9\n10,11,12\n",
		},
		{
			name: "Single row",
			input: [][]string{
				{"10", "20", "30"},
			},
			expected: "10,20,30\n",
		},
		{
			name:     "Empty matrix",
			input:    [][]string{},
			expected: "",
		},
		{
			name: "Empty rows",
			input: [][]string{
				{},
				{},
			},
			expected: "\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := m.Echo(tt.input)
			if output != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, output)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	m := NewCSVMatrixer()

	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name: "square matrix",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			expected: [][]string{
				{"1", "4", "7"},
				{"2", "5", "8"},
				{"3", "6", "9"},
			},
		},
		{
			name: "rectangular matrix",
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
				{"5", "6"},
			},
			expected: [][]string{
				{"1", "3", "5"},
				{"2", "4", "6"},
			},
		},
		{
			name: "single column matrix",
			input: [][]string{
				{"1"},
				{"3"},
				{"5"},
			},
			expected: [][]string{
				{"1", "3", "5"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := m.Invert(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, output)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	m := NewCSVMatrixer()

	tests := []struct {
		name     string
		input    [][]string
		expected []string
	}{
		{
			name: "Multiple rows",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := m.Flatten(tt.input)

			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, output)
			}
		})
	}
}
