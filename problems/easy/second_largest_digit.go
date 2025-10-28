package easy

import (
	"unicode"
)

func secondHighest(s string) int {
	highest := -1
	secondHighest := -1

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			currentNumber := int(ch) - 48

			if currentNumber > highest {
				secondHighest = highest
				highest = currentNumber
			} else if currentNumber > secondHighest && currentNumber < highest {
				secondHighest = currentNumber
			}
		}
	}

	return secondHighest
}
