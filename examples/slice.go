package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("New C: ", c)
	t := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(t)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("sl2: ", s[2:5])
	fmt.Println(s[:5])
	fmt.Println(s[2:])

	twoD := make([][]int, 3)
	for i := 0; i<3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D: ", twoD)
}
