package cpf

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"math/rand"

	"github.com/Hublastt/ValidGen/cmd/cpf/generate"
	"github.com/Hublastt/ValidGen/cmd/cpf/validade"
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

func (CPFFunctions) ValidateCPF(cpf string) error {
	cpfDigits := validade.SanitizeCPF(cpf)

	if len(cpfDigits) != 11 {
		return ErrInvalidLen
	}

	if validade.AllDigitsAreEqual(cpfDigits) {
		return ErrAllSameDigits
	}

	if !validade.ValidateVerificationDigits(cpfDigits) {
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

	numbers = append(numbers, generate.CalculateCheckDigit(numbers, 10))
	numbers = append(numbers, generate.CalculateCheckDigit(numbers, 11))

	cpf := ""

	for _, num := range numbers {
		cpf += fmt.Sprintf("%d", num)
	}

	if err := (CPFFunctions{}).ValidateCPF(cpf); err != nil {
		return "", err
	}

	return cpf, nil
}

func (CPFFunctions) FormatCPF(cpf string) (string, error) {

	if err := (CPFFunctions{}).ValidateCPF(cpf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s.%s-%s", cpf[:3], cpf[3:6], cpf[6:9], cpf[9:]), nil
}

func (CPFFunctions) UnformatCPF(cpf string) (string, error) {

	if err := (CPFFunctions{}).ValidateCPF(cpf); err != nil {
		return "", err
	}

	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	return cpf, nil
}
