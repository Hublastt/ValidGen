package cpf

import "testing"

func TestAllDigitsAreEqual(t *testing.T) {
	tests := []struct {
		cpf      string
		expected bool
	}{
		{"11111111111", true},
		{"22222222222", true},
		{"33333333333", true},
		{"12345678901", false},
		{"00000000000", true},
		{"", true},
		{"1", true},
		{"12", false},
	}

	for _, test := range tests {
		result := AllDigitsAreEqual(test.cpf)
		if result != test.expected {
			t.Errorf("AllDigitsAreEqual(%q) = %v; want %v", test.cpf, result, test.expected)
		}
	}
}
