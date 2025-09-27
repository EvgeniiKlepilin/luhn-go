# luhn-go

[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/EvgeniiKlepilin/luhn-go)](https://goreportcard.com/report/github.com/EvgeniiKlepilin/luhn-go)
[![Coverage](https://img.shields.io/badge/coverage-50%25-brightgreen.svg)](https://github.com/EvgeniiKlepilin/blockchain-go)

A fast and efficient implementation of the Luhn Algorithm in Go for validating credit card numbers and calculating check digits.

## What is the Luhn Algorithm?

The Luhn algorithm, also known as the "modulus 10" algorithm, is a simple checksum formula used to validate various identification numbers, including credit card numbers, IMEI numbers, and Canadian Social Insurance Numbers. It was created by IBM researcher Hans Peter Luhn and is widely used in the financial industry.

## Features

- ✅ **Validate credit card numbers** using the Luhn algorithm
- ✅ **Generate check digits** for incomplete numbers
- ✅ **Command-line interface** for easy usage
- ✅ **Interactive mode** for single number validation
- ✅ **Batch processing** for multiple numbers
- ✅ **Comprehensive error handling** for invalid inputs

## Supported Card Types

The implementation works with all major credit card types that use the Luhn algorithm:

- Visa
- MasterCard
- American Express
- Discover
- JCB
- Diners Club
- And more...

## Installation

### Prerequisites

- Go 1.24.2 or later

### Install from source

```bash
git clone https://github.com/EvgeniiKlepilin/luhn-go.git
cd luhn-go
go build -o luhn
```

### Install directly with go install

```bash
go install github.com/EvgeniiKlepilin/luhn-go@latest
```

## Usage

### Command Line Interface

#### Interactive Mode

Run the program without arguments to enter interactive mode:

```bash
./luhn
```

You'll be prompted to enter a number:

```
Please enter a number to validate:
4826493525582251
Validating 4826493525582251
Number 4826493525582251 is valid.
```

#### Batch Mode

Validate multiple numbers by passing them as command-line arguments:

```bash
./luhn 4826493525582251 373068810980009 5412769127028265
```

Output:
```
Validating the following numbers: [4826493525582251 373068810980009 5412769127028265]
Number 4826493525582251 is valid.
Number 373068810980009 is valid.
Number 5412769127028265 is valid.
```

#### Single Number Validation

```bash
./luhn 4826493525582251
```

### As a Go Library

You can also use this package as a library in your Go applications:

```go
package main

import (
    "fmt"
    "github.com/EvgeniiKlepilin/luhn-go"
)

func main() {
    // Validate a credit card number
    valid, err := isValid("4826493525582251")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    if valid {
        fmt.Println("Card number is valid!")
    } else {
        fmt.Println("Card number is invalid.")
    }
    
    // Generate a check digit
    checkDigit, err := getCheckDigit("482649352558225")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Check digit: %s\n", checkDigit) // Output: Check digit: 1
}
```

## API Reference

### Functions

#### `isValid(num string) (bool, error)`

Validates whether a given number is valid according to the Luhn algorithm.

**Parameters:**
- `num`: The number to validate as a string

**Returns:**
- `bool`: `true` if the number is valid, `false` otherwise
- `error`: An error if the input contains non-numeric characters

#### `getCheckDigit(num string) (string, error)`

Calculates the check digit needed to make a number valid according to the Luhn algorithm.

**Parameters:**
- `num`: The number (without check digit) as a string

**Returns:**
- `string`: The calculated check digit
- `error`: An error if the input contains non-numeric characters

## Examples

### Valid Credit Card Numbers

```bash
./luhn 4826493525582251    # Valid Visa
./luhn 5412769127028265    # Valid MasterCard  
./luhn 373068810980009     # Valid American Express
```

### Invalid Credit Card Numbers

```bash
./luhn 4826493525582257    # Invalid Visa
./luhn 5412769127028264    # Invalid MasterCard
./luhn 373068810980001     # Invalid American Express
```

### Error Cases

```bash
./luhn "1234 5678"         # Error: contains spaces
./luhn "1234-5678"         # Error: contains hyphens
./luhn "123a456"           # Error: contains letters
```

## Development

### Running Tests

```bash
go test -v
```

### Running Benchmarks

```bash
go test -bench=.
```

### Test Coverage

```bash
go test -cover
```

## Algorithm Implementation

The implementation follows the standard Luhn algorithm pseudocode from [Wikipedia](https://en.wikipedia.org/wiki/Luhn_algorithm):

1. Starting from the rightmost digit (excluding check digit) and moving left, double the value of every second digit
2. If the result of doubling is greater than 9, subtract 9 from it
3. Sum all the digits
4. The check digit is the amount needed to make the total sum a multiple of 10

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Evgenii Pendragon**

## Acknowledgments

- Hans Peter Luhn for creating the original algorithm
- The Go community for excellent tooling and documentation
- Wikipedia for the clear algorithm specification
