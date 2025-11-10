package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func main() {
	var inputTemp, scale, numPart, strPart string
	var result, num float64

	fmt.Print("Enter the temperature eg. 10K, 12F, 14C: ")
 	fmt.Scan(&inputTemp)
	fmt.Print("Enter the scale to convert eg. K, C, F: ")
	fmt.Scan(&scale)
	
	for _, ch := range inputTemp {
		if unicode.IsDigit(ch) || ch =='.' {
			numPart += string(ch)
		} else {
			strPart += string(ch)
		}
	}

	num, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Number:", numPart)
	fmt.Println("Letter:", strPart)

	if strPart == "K" && scale == "C" {
		result = num - 273.15
	} else if strPart == "K" && scale == "F" {
		result = (num -273.15) * 9/5 + 32
	} else if strPart == "C" && scale == "F" {
		result = (num * 9/5) + 32
	} else if strPart == "C" && scale == "K" {
		result = num + 273.15
	} else if strPart == "F" && scale == "C" {
		result = (num - 32) * 5/9
	} else if strPart == "F" && scale == "K" {
		result = (num - 32) * 5/9 + 273.15
	} else {
		fmt.Println("Unknown scale. Try Again.")
		return
	}

	fmt.Println("The result after conversion is:", result)
}



