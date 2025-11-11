package main

import( "fmt"
	"math/rand"
	"time"
)

func main () {
	
	nums := [10] int {1, 2, 3, 4, 5, 6, 7 ,8, 9, 10}
	
	rand.Seed(time.Now().UnixNano()) // ensures different results each run.	
	
	// Fisher-Yates Shuffle

	for i := len(nums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	
	var input int

	for _, v := range nums {
		fmt.Print("Guess a number between 1 - 10: ")
		fmt.Scan(&input)

		if (input == v) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}





}
