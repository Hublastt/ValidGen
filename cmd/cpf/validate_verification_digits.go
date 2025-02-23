package cpf

import "strconv"

func ValidateVerificationDigits(cpf string) bool {

	if cpf == "00000000000" {
		return false
	}

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
