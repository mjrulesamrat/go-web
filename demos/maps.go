package main

import "fmt"

func main() {
	elements := make(map[string]int)
	elements["H"] = 1
	fmt.Println(elements["H"])

	elements["O"] = 8
	delete(elements, "O")
	fmt.Println(elements["O"])

	if number, ok := elements["O"]; ok {
		fmt.Println(number)
	}
	if number, ok := elements["H"]; ok {
		fmt.Println(number)
	}
}
