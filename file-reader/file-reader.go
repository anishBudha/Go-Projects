package main

import (
	"fmt"
	"os"
	"bufio"
)


func main () {
	var filename string
	fmt.Println("Enter the name of the file to read")
	fmt.Print(": ")
	fmt.Scan(&filename)

	readWholeFile(filename)
	readLineByLine(filename)

}

func readWholeFile(filename string) {
	fmt.Println("\n Reading full file... \n")
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func readLineByLine(filename string) {
	fmt.Println("\n Reading line by line... \n")
	
	data, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
