package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)
	numbers[2] = 100

	some_numbers := numbers[1:3]
	fmt.Println(some_numbers)
	fmt.Println(len(numbers))
	
	var scores []float64
	scores = append(scores, 1.1)
	fmt.Println(scores)

	scores[0] = 2.2
	fmt.Println(scores)
}
