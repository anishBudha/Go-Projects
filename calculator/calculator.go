package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ") // This prints without a newline while Println prints with a new line
	fmt.Scanln(&num1)                 // & will pass the address of variable num1 to the function which will save the input in that adderess

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Cannot divide by zero")
			return // empty return will close the program
		}
	default:
		fmt.Println("Invalid operator")
		return
	}
	fmt.Println("Result:", result)
}
