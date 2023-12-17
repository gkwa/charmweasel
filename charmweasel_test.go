// charmweasel_test.go
package charmweasel

import (
	"testing"
)

func TestVariadicFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    []CustomType
		expected string
	}{
		{
			name:     "TestFooBarBaz",
			input:    []CustomType{Foo, Bar, Baz},
			expected: "Foo\nBar\nBaz\n",
		},
		{
			name:     "TestEmpty",
			input:    []CustomType{},
			expected: "",
		},
		{
			name:     "TestFoo",
			input:    []CustomType{Foo},
			expected: "Foo\n",
		},
		{
			name:     "TestInt",
			input:    []CustomType{42}, // 42 represents an unknown CustomType
			expected: "Unknown\n",
		},
		{
			name:     "TestNegative",
			input:    []CustomType{-42}, // -42 represents an unknown CustomType
			expected: "Unknown\n",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := VariadicFunction(tt.input...)
			if output.String() != tt.expected {
				t.Errorf("Expected output:\n%s\nActual output:\n%s\n", tt.expected, output.String())
			}
		})
	}
}
