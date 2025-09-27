package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		hasError bool
	}{
		// Valid credit card numbers
		{
			name:     "Valid American Express",
			input:    "373068810980009",
			expected: true,
			hasError: false,
		},
		{
			name:     "Valid MasterCard",
			input:    "5412769127028265",
			expected: true,
			hasError: false,
		},
		{
			name:     "Valid Visa",
			input:    "4826493525582251",
			expected: true,
			hasError: false,
		},
		// Invalid credit card numbers
		{
			name:     "Invalid American Express",
			input:    "373068810980001",
			expected: false,
			hasError: false,
		},
		{
			name:     "Invalid MasterCard",
			input:    "5412769127028264",
			expected: false,
			hasError: false,
		},
		{
			name:     "Invalid Visa",
			input:    "4826493525582257",
			expected: false,
			hasError: false,
		},
		// Error cases
		{
			name:     "Empty string",
			input:    "",
			expected: false,
			hasError: false,
		},
		{
			name:     "Non-numeric characters",
			input:    "123a456",
			expected: false,
			hasError: true,
		},
		{
			name:     "Contains spaces",
			input:    "1234 5678",
			expected: false,
			hasError: true,
		},
		{
			name:     "Contains hyphens",
			input:    "1234-5678",
			expected: false,
			hasError: true,
		},
		{
			name:     "Mixed alphanumeric",
			input:    "4826abc525582251",
			expected: false,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := isValid(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input %q, but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("isValid(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetCheckDigit(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		// Valid inputs - these should return the check digit needed to make the number valid
		{
			name:     "American Express without check digit",
			input:    "37306881098000",
			expected: "9",
			hasError: false,
		},
		{
			name:     "MasterCard without check digit",
			input:    "541276912702826",
			expected: "5",
			hasError: false,
		},
		{
			name:     "Visa without check digit",
			input:    "482649352558225",
			expected: "1",
			hasError: false,
		},
		{
			name:     "Two digits",
			input:    "12",
			expected: "5",
			hasError: false,
		},
		{
			name:     "Simple case",
			input:    "123",
			expected: "0",
			hasError: false,
		},
		// Edge cases that should produce check digit 0
		{
			name:     "Check digit should be 0",
			input:    "79927398712",
			expected: "0",
			hasError: false,
		},
		// Error cases
		{
			name:     "Empty string",
			input:    "",
			expected: "",
			hasError: false,
		},
		{
			name:     "Non-numeric characters",
			input:    "123a456",
			expected: "",
			hasError: true,
		},
		{
			name:     "Contains spaces",
			input:    "1234 567",
			expected: "",
			hasError: true,
		},
		{
			name:     "Contains special characters",
			input:    "1234-567",
			expected: "",
			hasError: true,
		},
		{
			name:     "Mixed characters",
			input:    "12ab34cd",
			expected: "",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getCheckDigit(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input %q, but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("getCheckDigit(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Integration test to verify that getCheckDigit works correctly with isValid
func TestGetCheckDigitIntegration(t *testing.T) {
	testCases := []string{
		"37306881098000",  // American Express without check digit
		"541276912702826", // MasterCard without check digit
		"482649352558225", // Visa without check digit
		"123456789",       // Random number
		"12",              // Two digits
	}

	for _, tc := range testCases {
		t.Run("Integration_"+tc, func(t *testing.T) {
			checkDigit, err := getCheckDigit(tc)
			if err != nil {
				t.Fatalf("getCheckDigit(%q) returned error: %v", tc, err)
			}

			// Append the check digit to create a complete number
			completeNumber := tc + checkDigit

			// Verify that the complete number is valid
			isValidResult, err := isValid(completeNumber)
			if err != nil {
				t.Fatalf("isValid(%q) returned error: %v", completeNumber, err)
			}

			if !isValidResult {
				t.Errorf("Number %q with check digit %q should be valid, but isValid returned false", tc, checkDigit)
			}
		})
	}
}

// Benchmark tests
func BenchmarkIsValid(b *testing.B) {
	testNumber := "4826493525582251"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = isValid(testNumber)
	}
}

func BenchmarkGetCheckDigit(b *testing.B) {
	testNumber := "482649352558225"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = getCheckDigit(testNumber)
	}
}
