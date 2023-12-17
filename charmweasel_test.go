package charmweasel

import (
	"log/slog"
	"testing"
)

func TestVariadicFunction(t *testing.T) {
	tests := []struct {
		name     string
		level    slog.Level
		input    []CustomType
		expected string
	}{
		{
			name:     "TestFooBarBaz",
			level:    slog.LevelInfo,
			input:    []CustomType{Foo, Bar, Baz},
			expected: "Foo\nBar\nBaz\nLog level: INFO\n",
		},
		{
			name:     "TestEmpty",
			level:    slog.LevelError,
			input:    []CustomType{},
			expected: "Log level: ERROR\n",
		},
		{
			name:     "TestFoo",
			level:    slog.LevelDebug,
			input:    []CustomType{Foo},
			expected: "Foo\nLog level: DEBUG\n",
		},
		{
			name:     "TestInt",
			level:    slog.LevelWarn,
			input:    []CustomType{42}, // 42 represents an unknown CustomType
			expected: "Unknown\nLog level: WARN\n",
		},
		{
			name:     "TestNegative",
			level:    slog.LevelError,
			input:    []CustomType{-42}, // -42 represents an unknown CustomType
			expected: "Unknown\nLog level: ERROR\n",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := VariadicFunction(tt.level, tt.input...)
			if output.String() != tt.expected {
				t.Errorf("Expected output:\n%s\nActual output:\n%s\n", tt.expected, output.String())
			}
		})
	}
}
