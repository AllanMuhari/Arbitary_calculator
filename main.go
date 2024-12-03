package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Arbitrary Precision Calculator")
	fmt.Println("Input your calculations (e.g., '12345 + 67890') or type 'exit' to exit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		result := processInput(input)
		fmt.Printf("Result: %s\n", result)
	}
}

func processInput(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "Invalid input. Please provide two numbers and an operator."
	}

	number1 := NewBigNumber(parts[0], 10)
	number2 := NewBigNumber(parts[2], 10)
	operator := parts[1]

	switch operator {
	case "+":
		result := Add(number1, number2)
		return result.String()
	case "-":
		result := Subtract(number1, number2)
		return result.String()
	case "*":
		result := Multiply(number1, number2)
		return result.String()
	case "/":
		quotient, remainder := Divide(number1, number2)
		return fmt.Sprintf("Quotient: %s, Remainder: %s", quotient.String(), remainder.String())
	default:
		return "Invalid operator. Please use +, -, *, or /."
	}
}
