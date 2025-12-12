package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main () {
	fmt.Println("=== BMI Calculator ===")
	fmt.Println("Body Mass Index Calculator")
	fmt.Println()

	weight := getFloatInput("Enter your weight in kilograms: ")
	height := getFloatInput("Enter your height in meters: ")
	bmi := calculateBMI(weight, height)
	category := getBMICategory(bmi)
	
	fmt.Println()
	fmt.Printf("Your BMI: %.2f\n", bmi)
	fmt.Printf("Category: %s\n", category)

}

func getFloatInput(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64) 
		if value <= 0 {
			fmt.Println("Value must be greater than zero.")
			continue
		}

		return value
	}
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func getBMICategory(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi < 25 {
		return "Overweight"
	} else {
		return "Obese"
	}
}
