package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Enter:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)
}
