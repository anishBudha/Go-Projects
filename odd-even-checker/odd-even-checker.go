package main

import(
	"fmt"
	"strconv"
)

func main () {
	var input string
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	
	num, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid, Please input only integer.")
		return
	}

	remainder := num % 2

	if remainder == 0 {
		fmt.Println(input, "is a even number")
	} else if remainder != 0 {
		fmt.Println(input, "is an odd number")
	}
}
