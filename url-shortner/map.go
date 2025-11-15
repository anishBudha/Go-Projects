package main

import "fmt"

func main () {
	myMap := make(map[string]int) // creating an empty map 
	myMap["age"] = 20
	
	newMap := make(map[string]string)
	newMap["name"] = "Anish"

	againMap := make(map[string]bool)
	againMap["isMarried"] = false

	aSlice := make([]int, 3)
	aSlice[0] = 0
	aSlice[1] = 1
	aSlice[2] = 2

	fmt.Println(myMap["age"])
	fmt.Println(newMap["name"])
	fmt.Println(againMap["isMarried"])
	fmt.Println(aSlice)
}
