package validation

import (
	`strconv`
)

const (
	iinLength = 12
)

var (
	w1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	w2 = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}
)

func ValidateIdNumber(iin string) bool {
	if len(iin) != iinLength {
		return false
	}

	checkSum := multiply(iin[:len(iin)-1], w1) % 11
	if checkSum == 10 {
		checkSum = multiply(iin[:len(iin)-1], w2) % 11
	}

	lastDigit, _ := strconv.Atoi(string(iin[11]))
	return checkSum == lastDigit
}

func multiply(iin string, w []int) int {
	if len(iin) != len(w) {
		return 0
	}
	result := 0
	for i, c := range iin {
		digit, _ := strconv.Atoi(string(c))
		result += digit * w[i]
	}
	return result
}
