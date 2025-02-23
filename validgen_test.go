package validgen

import (
	"testing"
)

func TestGenerateCpf(t *testing.T) {
	cpf, err := GenerateCpf()
	if err != nil {
		t.Errorf("GenerateCpf() error = %v", err)
	}
	if len(cpf) != 11 {
		t.Errorf("GenerateCpf() = %v; want length 11", cpf)
	}
}

func TestValidateCpf(t *testing.T) {
	validCpf := "12345678909"
	invalidCpf := "12345678900"

	if err := ValidateCpf(validCpf); err != nil {
		t.Errorf("ValidateCpf(%v) error = %v; want nil", validCpf, err)
	}

	if err := ValidateCpf(invalidCpf); err == nil {
		t.Errorf("ValidateCpf(%v) error = nil; want error", invalidCpf)
	}
}
