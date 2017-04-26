package main

import "fmt"

// Buffered channels accept the
// specified number of values before they are received

func main() {
	messages := make(chan string, 2)
	messages <- "Golang"
	messages <- "Gopher"
	
	//Recieve value from buffered channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
