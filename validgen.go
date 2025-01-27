package validgen

import "github.com/Hublastt/ValidGen/cmd/cpf"

func GenerateCpf() (string, error) {
	return cpf.CPFFunctions{}.GerarCPF()
}

func ValidateCpf(cpfString string) error {
	return cpf.CPFFunctions{}.IsValid(cpfString)
}

func ValidGen() cpf.ValidGenStruct {
	return cpf.ValidGenStruct{
		CPF: cpf.CPFFunctions{},
	}
}
