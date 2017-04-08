package main

import "fmt"

func main() {
x := 1
if x != 0 {
    fmt.Println("Yes")
}
var y []string
fmt.Println(len(y))
if len(y) != 0 {
    fmt.Println("this won't be printed")
}
}
