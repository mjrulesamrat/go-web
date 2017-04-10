package main

import "fmt"

func main() {
	fmt.Printf("Hello, %s\n", new(World))
}

type World struct{}

func (w *World) String() string {
	return "foodies of ahmedabad."
}
