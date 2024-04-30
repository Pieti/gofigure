package luhn

import "testing"

func TestCalculateChecksum(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1789372997", 4},
	}

	for _, test := range tests {
		if got := CalculateChecksum(test.input); got != test.expected {
			t.Errorf("CalculateChecksum(%s) = %d, want %d", test.input, got, test.expected)
		}
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"79927398713", true},
		{"7992 7398 713", true},
		{"79927398714", false},
		{"", false},
		{"abcdefghijk", false},
	}

	for _, test := range tests {
		if got := Validate(test.input); got != test.expected {
			t.Errorf("Validate(%s) = %t, want %t", test.input, got, test.expected)
		}
	}
}
