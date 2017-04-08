package main

import "fmt"

func main() {
	names := []string{"jay", "krishna", "sam"}

	for i, name := range names {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}
