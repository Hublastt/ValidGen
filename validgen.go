package validgen

import "github.com/Hublastt/ValidGen/cmd/cpf"

/*
ValidateCpf - Validate cpf is valid
  - @param {string} - cpf
  - @return {error} - error
*/
func ValidateCPF(cpfString string) error {
	return cpf.CPFFunctions{}.ValidateCPF(cpfString)
}

/*
GenerateCpf - Generate a valid cpf
  - @return {string, error} - cpf, error
*/
func GenerateCPF() (string, error) {
	return cpf.CPFFunctions{}.GenerateCPF()
}

/*
FormatCpf - Format cpf
  - @param {string} - cpf
  - @return {string, error} - cpf, error
*/
func FormatCPF(cpfString string) (string, error) {
	return cpf.CPFFunctions{}.FormatCPF(cpfString)
}

/*
UnformatCpf - Unformat cpf
  - @param {string} - cpf
  - @return {string, error} - cpf, error
*/
func UnformatCPF(cpfString string) (string, error) {
	return cpf.CPFFunctions{}.UnformatCPF(cpfString)
}
