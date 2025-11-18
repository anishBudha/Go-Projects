package main

import (
	"fmt"
	"os"
	"encoding/json"	
)

type JsonData struct {
	Name string `json:"name"`
	Net string `json:"net"`
	Allocated bool `json:"allocated"`
	Allocation int `json:"allocation"`
}

func main() {
	output := readfile()
	
	for _, item := range output {
		fmt.Println("Name:", item.Name)
		fmt.Println("Net:", item.Net)
		fmt.Println("Allocated:", item.Allocated)
		fmt.Println("Allocation:", item.Allocation)
		fmt.Print("\n")
	}

}

func readfile() []JsonData {
	fmt.Print("\n")
	fmt.Print("Enter the json filename: ")
	var filename string
	fmt.Scan(&filename)
	fmt.Print("\n")

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var jsonExt []JsonData
	err = json.Unmarshal(data, &jsonExt)

	return jsonExt
}
