package main

import (
	"strconv"
)

// Pseudocode for Luhn Algorithm
// Source: https://en.wikipedia.org/wiki/Luhn_algorithm
//
// function isValid(cardNumber[1..length])
//     sum := 0
//     parity := length mod 2
//     for i from 1 to (length - 1) do
//         if i mod 2 == parity then
//             sum := sum + cardNumber[i]
//         elseif cardNumber[i] > 4 then
//             sum := sum + 2 * cardNumber[i] - 9
//         else
//             sum := sum + 2 * cardNumber[i]
//         end if
//     end for
//     return cardNumber[length] == ((10 - (sum mod 10)) mod 10)
// end function

func isValid(num string) (bool, error) {
	if len(num) < 1 {
		return false, nil
	}
	checkDigit, err := getCheckDigit(num[0 : len(num)-1])
	if err != nil {
		return false, err
	}
	return string(num[len(num)-1]) == checkDigit, nil
}

func getCheckDigit(num string) (string, error) {
	if len(num) < 1 {
		return "", nil
	}
	sum := 0
	parity := len(num) % 2
	for i := range len(num) {
		digit, err := strconv.Atoi(string(num[i]))
		if err != nil {
			return "", err
		}
		if i%2 == parity {
			sum += digit
		} else if digit > 4 {
			sum += 2*digit - 9
		} else {
			sum += 2 * digit
		}
	}
	return strconv.Itoa((10 - (sum % 10)) % 10), nil
}
