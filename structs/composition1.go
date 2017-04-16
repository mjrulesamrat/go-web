// Example program for composotion of types

package main

import (
	"fmt"
	_ "time"
)

type Person struct {
	FirstName, LastName string
	Email, Location string
}

func (p Person) PrintName() {
	fmt.Printf("%s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[Email : %s, Location:  %s ]\n", p.Email, p.Location)
}

type Admin struct {
	Person //type embedding fo composition
	Roles []string
}

func (a Admin) PrintDetails() {
	a.Person.PrintDetails()
	fmt.Println("Admin Roles: ")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

type Member struct {
	Person // type embadding for composition
	Skills []string
}

func (m Member) PrintDetails() {
	m.Person.PrintDetails()
	fmt.Println("Member Skills: ")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}


func main() {

	jay := Member{
			Person{
				"Jay",
				"Modi",
				"mjrulesamrat@gmail.com",
				"Ahmedabad",
			},
			[]string{"Python", "Go", "Docker", "Kubernetes"},
	}

	krishna := Admin{
			Person{
				"Krishna",
				"Yadav",
				"krishna@example.com",
				"Dwarka",
			},
			[]string{"Manage World", "Manage Karma"},
	}

	// call method for jay
	jay.PrintName()
	jay.PrintDetails()

	// call method for krishna
	krishna.PrintName()
	krishna.PrintDetails()
}