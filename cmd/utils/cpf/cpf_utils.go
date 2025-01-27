package utils

import (
	"regexp"
	"strconv"
)

func SanitizeCPF(cpf string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(cpf, "")
}

func AllDigitsAreEqual(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func ValidateVerificationDigits(cpf string) bool {
	for i := 9; i < 11; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			digit, _ := strconv.Atoi(string(cpf[j]))
			sum += digit * (i + 1 - j)
		}

		verificationDigit := (sum * 10) % 11
		if verificationDigit == 10 {
			verificationDigit = 0
		}

		if verificationDigit != int(cpf[i]-'0') {
			return false
		}
	}
	return true
}

func CalcularDigitoVerificador(numeros []int, pesoInicial int) int {
	soma := 0
	peso := pesoInicial

	for _, num := range numeros {
		soma += num * peso
		peso--
	}

	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}
