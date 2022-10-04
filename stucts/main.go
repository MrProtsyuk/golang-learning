package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}


//use this type if you want to access "0 value" of a variable
// func main() {
// 	var mark person

// 	mark.firstName = "Mark"
// 	mark.lastName = "Protsyuk"

// this just shows value
// fmt.Println(mark)
// }

// syntax here is not standard but it gives more leaniancy with variables
func main() {
	mark := person{
		firstName: "Mark", 
		lastName: "Protsyuk",
		// embedded stuct
		contactInfo: contactInfo{
			email: "Mark@gmail.com",
			zipCode: 12345,
		},
	}

	//ampersand gives us access to RAM address of the variable
	// markPointer := &mark
	mark.updateName("Marky")
	mark.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	// This shows value and variable
	// can be written with or without second argument; fmt.Printf ("%v+", p)
  fmt.Println(p)
}
