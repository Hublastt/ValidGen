package validgen

import (
	"testing"
)

func TestGenerateCpf(t *testing.T) {
	cpf, err := GenerateCPF()
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

	if err := ValidateCPF(validCpf); err != nil {
		t.Errorf("ValidateCpf(%v) error = %v; want nil", validCpf, err)
	}

	if err := ValidateCPF(invalidCpf); err == nil {
		t.Errorf("ValidateCpf(%v) error = nil; want error", invalidCpf)
	}
}

func TestFormatGeneratedCpf(t *testing.T) {
	formattedCpf, err := FormatGeneratedCPF()
	if err != nil {
		t.Errorf("FormatGeneratedCpf() error = %v", err)
	}
	if len(formattedCpf) != 14 {
		t.Errorf("FormatGeneratedCpf() = %v; want length 14", formattedCpf)
	}
	if formattedCpf[3] != '.' || formattedCpf[7] != '.' || formattedCpf[11] != '-' {
		t.Errorf("FormatGeneratedCpf() = %v; want format XXX.XXX.XXX-XX", formattedCpf)
	}
}
