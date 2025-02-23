package cpf

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name    string
		cpf     string
		wantErr error
	}{
		{"Valid CPF", "12345678909", nil},
		{"Invalid length", "1234567890", ErrInvalidLen},
		{"All same digits", "11121111111", ErrAllSameDigits},
		{"Invalid CPF", "12345678900", ErrInvalidCPF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := (CPFFunctions{}).ValidateCPF(tt.cpf)
			if err != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGerarCPF(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run("Generate CPF", func(t *testing.T) {
			cpf, err := (CPFFunctions{}).GenerateCPF()
			if err != nil {
				t.Errorf("GerarCPF() error = %v", err)
			}
			if err := (CPFFunctions{}).ValidateCPF(cpf); err != nil {
				t.Errorf("Generated CPF is invalid: %v", err)
			}
		})
	}
}

func TestUnFortmatCPF(t *testing.T) {
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
		result, _ := (CPFFunctions{}).UnformatCPF(test.input)
		if result != test.expected {
			t.Errorf("SanitizeCPF(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
func TestFormatCPF(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  error
	}{
		{"12345678909", "123.456.789-09", nil},
		{"1234567890", "", ErrInvalidLen},
		{"11111111111", "", ErrAllSameDigits},
		{"12345678900", "", ErrInvalidCPF},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := (CPFFunctions{}).FormatCPF(test.input)
			if err != test.wantErr {
				t.Errorf("FormatCPF() error = %v, wantErr %v", err, test.wantErr)
			}
			if result != test.expected {
				t.Errorf("FormatCPF(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}
