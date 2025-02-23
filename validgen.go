package validgen

import "github.com/Hublastt/ValidGen/cmd/cpf"

/*
GenerateCpf - Generate a valid cpf
  - @return {string, error} - cpf, error
*/
func GenerateCpf() (string, error) {
	return cpf.CPFFunctions{}.GenerateCPF()
}

/*
ValidateCpf - Validate cpf is valid
  - @param {string} - cpf
  - @return {error} - error
*/
func ValidateCpf(cpfString string) error {
	return cpf.CPFFunctions{}.IsValid(cpfString)
}
