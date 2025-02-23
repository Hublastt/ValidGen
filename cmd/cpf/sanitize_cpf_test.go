package cpf

import "testing"

func TestSanitizeCPF(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"123.456.789-09", "12345678909"},
		{"123 456 789 09", "12345678909"},
		{"123-456-789-09", "12345678909"},
		{"12345678909", "12345678909"},
		{"abc123def456ghi789jkl09", "12345678909"},
	}

	for _, test := range tests {
		result := SanitizeCPF(test.input)
		if result != test.expected {
			t.Errorf("SanitizeCPF(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
