// Person struct with methods of pointer reciever
package main

import (
	"fmt"
	"time"
)

// Go does not provide classes in it's type system; it has structs, which are analogous to classes.

/* Structs in Go let you create user-defined concrete types; they are collections of fields 
   or properties that can be used for storing complex data. You can use structs for storing
   application domain models.

*/

// Declaring a Struct with a Group of Fields
type Person struct {
	FirstName, LastName    string
	Dob    time.Time
	Email, Location    string
}

// Adding Behavior to a Struct Type
// A person method
// here (p Person) is a receiver
func (p Person) PrintName() {
	fmt.Printf("%s %s\n", p.FirstName, p.LastName)
}

// you can also define pointer receiver too
// A person method with pointer receiver
func (p *Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

// A person method with dynamic changing location name. Possible only with pointer receiver methods.

func (p1 *Person) ChangeLocation(newLocation string) {
	p1.Location= newLocation
}

func main() {

	// creating Sturct Instance

	var p Person
	p.FirstName="Jay"
	p.LastName="Modi"
	p.Dob=time.Date(1992, time.November, 2, 0, 0, 0, 0, time.UTC)
	p.Email="mjrulesamrat@gmail.com"
	p.Location="Ahmedabad"

	p.PrintName()

	// creating using struct Literal
	/* Because you want to call a pointer receiver method, you can create the Person type instance by
		providing the ampersand (&) operator. */
	p1 := &Person{
		"Jay",
		"Modi",
		time.Date(1992, time.November, 2, 0, 0, 0, 0, time.UTC),
		"mjrulesamrat@gmail.com",
		"Ahmedabad",
	}
	
	p1.PrintDetails()
	p1.ChangeLocation("Mumbai")
	p1.PrintDetails()

}
