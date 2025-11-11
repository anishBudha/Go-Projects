//package main
//
//import "fmt"
//
//func main () {
//	for i:= 1; i <= 100; i++ {
//		var result string
//		if i % 2 == 0 {
//			result = "Fizz"
//			
//		}
//
//		if i % 5 == 0 {
//			result += "Buzz"	
//		} 	
//
//		if result == "" {
//			fmt.Println(i)
//		} else {
//			fmt.Println(result)
//		}
//		
//	}
//}

package main

import "fmt"

func main () {
	for i:=1; i<=100; i++ {
		if i%2 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%2 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
