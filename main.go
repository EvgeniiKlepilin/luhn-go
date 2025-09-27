package main

import "fmt"

func main() {
	validCreditCardNumbers := []string{
		"373068810980006",
		"5412769127028262",
		"4826493525582251",
	}
	invalidCreditCardNumbers := []string{
		"373068810980001",
		"5412769127028264",
		"4826493525582257",
	}
	creditCardNumbers := append(validCreditCardNumbers, invalidCreditCardNumbers...)
	for _, num := range creditCardNumbers {
		valid, err := isValid(num)
		if err != nil {
			fmt.Println("Err:", err)
		} else if valid {
			fmt.Println("Credit card number", num, "is valid.")
		} else {
			fmt.Println("Credit card number", num, "is invalid.")
		}
	}
}
