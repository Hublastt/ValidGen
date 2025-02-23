package cpf

import "testing"

func TestValidateVerificationDigits(t *testing.T) {
	tests := []struct {
		cpf      string
		expected bool
	}{
		{"12345678909", true},
		{"11144477735", true},
		{"12345678900", false},
		{"11144477736", false},
		{"00000000000", false},
	}

	for _, test := range tests {
		t.Run(test.cpf, func(t *testing.T) {
			result := ValidateVerificationDigits(test.cpf)
			if result != test.expected {
				t.Errorf("ValidateVerificationDigits(%s) = %v; want %v", test.cpf, result, test.expected)
			}
		})
	}
}
