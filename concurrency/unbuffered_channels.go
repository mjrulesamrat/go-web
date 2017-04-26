package main

import "fmt"

func main() {

	count := make(chan int)

	go PrintCounts(count)
	fmt.Println("Channel begins")
	count <- 1
}

func PrintCounts(count chan int) {
	for {
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed \n")
			return
		}
		fmt.Printf("Count: %d received \n", val)
		if val == 10 {
			fmt.Println("Channel closed \n")
			close(count)
			return
		}
		val++
		count <- val
	}
}
