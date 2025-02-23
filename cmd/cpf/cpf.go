package cpf

import (
	"errors"
	"fmt"
	"time"

	"math/rand"
)

var (
	ErrInvalidLen    = errors.New("CPF must contain exactly 11 digits")
	ErrAllSameDigits = errors.New("invalid CPF: all digits are the same")
	ErrInvalidCPF    = errors.New("invalid CPF, please try again")
	ErrValidCPF      = errors.New("failed validation")
)

type CPFFunctions struct{}

type ValidGenStruct struct {
	CPF CPFFunctions
}

func (CPFFunctions) IsValid(cpf string) error {
	cpfDigits := SanitizeCPF(cpf)

	if len(cpfDigits) != 11 {
		return ErrInvalidLen
	}

	if AllDigitsAreEqual(cpfDigits) {
		return ErrAllSameDigits
	}

	if !ValidateVerificationDigits(cpfDigits) {
		return ErrInvalidCPF
	}

	return nil
}

func (CPFFunctions) GenerateCPF() (string, error) {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, 9)

	for i := range 9 {
		numbers[i] = source.Intn(10)
	}

	numbers = append(numbers, CalculateCheckDigit(numbers, 10))
	numbers = append(numbers, CalculateCheckDigit(numbers, 11))

	cpf := ""

	for _, num := range numbers {
		cpf += fmt.Sprintf("%d", num)
	}

	if err := (CPFFunctions{}).IsValid(cpf); err != nil {
		return "", ErrValidCPF
	}

	return cpf, nil
}
