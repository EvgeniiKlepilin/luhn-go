package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter a number to validate:")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println("Validating", input)
		valid, err := isValid(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else if valid {
			fmt.Println("Number", input, "is valid.")
		} else {
			fmt.Println("Number", input, "is invalid.")
		}
	} else {
		fmt.Println("Validating the following numbers:", args)
		for _, num := range args {
			valid, err := isValid(num)
			if err != nil {
				fmt.Println("Error:", err)
			} else if valid {
				fmt.Println("Number", num, "is valid.")
			} else {
				fmt.Println("Number", num, "is invalid.")
			}
		}
	}
}
