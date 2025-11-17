// package main
//
// import "fmt"
//
// func main () {
//
// 	// func functionName(paramName ...parameterType) returnType {
// 	// 	code block
// 	// 	return
// 	// }
// 	//
//
// 	sumInt := sum(1, 2, 3, 4, 5, 6)
//
// 	fmt.Println(sumInt)
//
//
// }
//
//
// func sum (nums ...int) int {
// 	total := 0
//
// 	for _, num := range nums {
// 		total += num
// 	}
//
// 	return total
// }


package main

import "fmt"


func main () {
	nums := []int{1, 2, 3, 4, 5}

	s := sum(nums...)

	fmt.Println(s)
}


func sum (nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// if needed to pass multiple values or optional values to the parameter, its better to use struct

type MathOperationOptions struct {
	Operation string
	ValueOne int
	ValueTwo int
}

func mathOperation(opts MathOperationOptions) {
	// code block
}
