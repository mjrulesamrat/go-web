package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func main() {

	count := make(chan int)

	wg.Add(2)
	fmt.Println("Start Goroutines")
	go PrintCounts(count, "A")
	go PrintCounts(count, "B")
	fmt.Println("Channel begins")
	count <- 1
	
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\n Termination Program")
}

func PrintCounts(count chan int, label string) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()
	for {
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed \n")
			return
		}
		fmt.Printf("Count: %d received from label %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel closed from label %s \n", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}
