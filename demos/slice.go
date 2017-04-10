package main

import "fmt"

func main() {
	x := []int{10, 20, 30}
	y := make([]int, 2)
	copy(y, x)
	fmt.Println(x, y)
}
