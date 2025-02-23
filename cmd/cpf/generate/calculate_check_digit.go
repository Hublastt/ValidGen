package generate

func CalculateCheckDigit(numbers []int, initialWeight int) int {
	sum := 0
	weight := initialWeight

	for _, num := range numbers {
		sum += num * weight
		weight--
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
