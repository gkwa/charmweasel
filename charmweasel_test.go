// charmweasel_test.go
package charmweasel

import (
	"bytes"
	"io"
	"os"
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
			name:     "TestUnknown",
			input:    []CustomType{42}, // 42 represents an unknown CustomType
			expected: "Unknown\n",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout to capture Println output
			output := captureOutput(func() {
				VariadicFunction(tt.input...)
			})

			// Check if the output matches the expected value
			if output != tt.expected {
				t.Errorf("Expected output:\n%s\nActual output:\n%s\n", tt.expected, output)
			}
		})
	}
}

// captureOutput captures the output of a function and returns it as a string.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()

	return buf.String()
}
