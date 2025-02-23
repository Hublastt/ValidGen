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
		{"All same digits", "11111111111", ErrAllSameDigits},
		{"Invalid CPF", "12345678900", ErrInvalidCPF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := (CPFFunctions{}).IsValid(tt.cpf)
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
			if err := (CPFFunctions{}).IsValid(cpf); err != nil {
				t.Errorf("Generated CPF is invalid: %v", err)
			}
		})
	}
}
