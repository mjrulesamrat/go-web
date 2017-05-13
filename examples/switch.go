package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a Bull :P")
		case int:
			fmt.Println("I'm an ant")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI("true")
	whatAmI(true)
	whatAmI(10)
}