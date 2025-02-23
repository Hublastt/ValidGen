package cpf

import "testing"

func TestCalculateCheckDigit(t *testing.T) {
	tests := []struct {
		numbers       []int
		initialWeight int
		expected      int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 10, 1},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 10, 0},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 10, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 11, 9},
	}

	for _, test := range tests {
		result := CalculateCheckDigit(test.numbers, test.initialWeight)
		if result != test.expected {
			t.Errorf("CalculateCheckDigit(%v, %d) = %v; want %v", test.numbers, test.initialWeight, result, test.expected)
		}
	}
}
