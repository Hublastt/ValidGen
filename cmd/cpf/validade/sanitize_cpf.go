package validade

import "regexp"

func SanitizeCPF(cpf string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(cpf, "")
}
