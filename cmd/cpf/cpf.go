package cpf

import (
	"errors"
	"fmt"
	"time"

	"math/rand"

	utils "github.com/Hublastt/ValidGen/cmd/utils/cpf"
)

var (
	errInvalidLen    = errors.New("cpf must contain exactly 11 digits")
	errAllSameDigits = errors.New("invalid CPF: all digits are the same")
	errInvalidCPF    = errors.New("invalid CPF, please try again")
	errValidCpf      = errors.New("failed validation")
)

type CPFFunctions struct{}

type ValidGenStruct struct {
	CPF CPFFunctions
}

func (CPFFunctions) IsValid(cpf string) error {
	cpfDigits := utils.SanitizeCPF(cpf)

	if len(cpfDigits) != 11 {
		return errInvalidLen
	}

	if utils.AllDigitsAreEqual(cpfDigits) {
		return errAllSameDigits
	}

	if !utils.ValidateVerificationDigits(cpfDigits) {
		return errInvalidCPF
	}

	return nil
}

func (CPFFunctions) GerarCPF() (string, error) {
	fonte := rand.New(rand.NewSource(time.Now().UnixNano()))
	numeros := make([]int, 9)

	for i := 0; i < 9; i++ {
		numeros[i] = fonte.Intn(10)
	}

	numeros = append(numeros, utils.CalcularDigitoVerificador(numeros, 10))
	numeros = append(numeros, utils.CalcularDigitoVerificador(numeros, 11))

	cpf := ""

	for _, num := range numeros {
		cpf += fmt.Sprintf("%d", num)
	}

	if err := (CPFFunctions{}).IsValid(cpf); err != nil {
		return "", errValidCpf
	}

	return cpf, nil
}
