package luhn

import "strings"

// Check if the provided card number is valid
func Validate(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	if len(cardNumber) < 2 {
		return false
	}

	for _, digit := range cardNumber {
		if digit < '0' || digit > '9' {
			return false
		}
	}

	checksum := CalculateChecksum(cardNumber[:len(cardNumber)-1])

	return checksum == int(cardNumber[len(cardNumber)-1]-'0')

}

// Calculate the Luhn checksum of the provided number
// Assume the number is a string of digits without checksum
func CalculateChecksum(number string) int {
	sum := 0
	doubled := true

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if doubled {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		doubled = !doubled
		sum += digit
	}
	return 10 - sum%10
}
